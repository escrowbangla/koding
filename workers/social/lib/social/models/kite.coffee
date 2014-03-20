jraphical = require 'jraphical'
module.exports = class JKite extends jraphical.Module

  KodingError         = require '../error'
  {permit}            = require './group/permissionset'
  {signature}         = require 'bongo'
  {Relationship}      = jraphical

  @trait __dirname, '../traits/protected'

  @share()
  @set
    permissions           :
      'create kite'       : ['member']
      'list kites'        : ['member']
    schema                :
      name                : String
      description         : String
      createdAt           :
        type              : Date
        default           : -> new Date
    sharedEvents          :
      instance            : []
      static              : []
    sharedMethods         :
      static              :
        create            :
          (signature Object, Function)
        list              :
          (signature Object, Object, Function)
        one               :
          (signature Object, Function)
      instance            :
        modify            :
          (signature Object, Function)
        deleteKite        :
          (signature Function)
        fetchPlans        :
          (signature Function)
        createPlan        :
          (signature Object, Function)
        deletePlan        :
          (signature Object, Function)

    relationships       :
      plan              :
        targetType      : 'JPaymentPlan'
        as              : 'kitePlan'

  @create: permit 'create kite',
    success:(client, formData, callback)->
      kite = new JKite formData
      kite.save (err)->
        return  callback new KodingError "kite couldn't saved" if err
        account = client.connection.delegate
        account.addKite kite, (err, res)->
          return  callback new KodingError "kite couldn't added to account" if err
          callback null, kite

  @list: permit 'list kites',
    success: (client, selector, options, callback)->
      JKite.some selector, options, callback

  modify: permit 'edit kite',
    success: (client, formData, callback)->
      @update $set: {name: formData.name, description: formData.description} , callback

  deleteKite: permit 'delete kite',
    success: (client, callback)->
      account = client.connection.delegate
      Relationship.one {
        targeName   : "JKite"
        targetId    : @getId()
        sourceName  : "JAccount"
        sourceId    : account.getId()
        as          : "owner"
      }, (err, rel) =>
        return  callback new KodingError err if err
        @remove (err, res)->
          return  callback err, res

