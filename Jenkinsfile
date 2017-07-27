 pipeline {
   agent any

   environment {
     GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.8.3"
   }

   stages {
     stage('Build') {
       steps {
         tool(name: 'Go 1.8.3', type: 'go')
         withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
             sh 'printenv'
             sh 'go version'
         }
         // sh 'go build'
       }
     }
   }
 }