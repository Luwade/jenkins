 pipeline {
   agent any

   environment {
     GOCONFIG_PATH="/home/vagrant/go"
   }

   stages {
     stage('Build') {
       steps {
         withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
             sh 'printenv'
             sh 'go version'
         }
         // sh 'go build'
       }
     }
   }
 }