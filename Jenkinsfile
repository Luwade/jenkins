 pipeline {
   agent any

   environment {
     GOCONFIG_PATH="/home/vagrant/go"
   }

   stages {
     stage('Build') {
       steps {
         withEnv(["GOROOT=$GOCONFIG_PATH", "GOPATH=$GOCONFIG_PATH"]) {
             sh 'printenv'
             sh 'go version'
         }
         // sh 'go build'
       }
     }
     stage('Test') {
       steps {
          withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
              sh 'go test'
          }
       }
     }
     stage('Deploy') {
        steps {
            withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
            withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'admin', usernameVariable: 'Luwade', passwordVariable: 'Luwade582']]) {
                sh('git push https://${GIT_USERNAME}:${GIT_PASSWORD}@<jenkins> ')
            }

               // sh 'git push origin master'
             }
         }
      }
   }
 }