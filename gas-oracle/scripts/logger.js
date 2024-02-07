const log4js =  require('log4js')
const path  = require('path')
const LOG_PATH = path.join('','logs')
 
log4js.configure({
    appenders:{
        logFile:{
            type:'dateFile',
            filename:path.join(LOG_PATH,"main"),
            pattern:'yyyy-MM-dd.log',
            keepFileExt: true,
            alwaysIncludePattern: true,
            layout:{
                type:'pattern',
                pattern:'%m%n'
            }
        },
        out:{
            type:'console'
        }
    },
    categories:{
        default:{
            appenders:['out'],
            level:'INFO'
        },
        main:{
            appenders:['out','logFile'],
            level:'DEBUG'
        }
    }
})
 
module.exports = {
    main : log4js.getLogger("main")
}