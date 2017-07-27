 pipeline {
   agent any

   environment {
     GOCONFIG_PATH="/home/vagrant/go"
   }

   stages {
     stage('Build') {
       steps {
         withEnv(["GOROOT=${root}", "GOPATH=$GOCONFIG_PATH"]) {
             sh 'printenv'
             sh 'go version'
         }
         // sh 'go build'
       }
     }
     stage('Test') {
        steps {
            withEnv(["GOROOT=${root}", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                sh 'go test'
            }
        }
     }
   }
 }