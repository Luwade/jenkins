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
                sh 'git remote add origin https://github.com/Luwade/jenkins.git'
                sh 'git push origin master'
             }
         }
      }
   }
 }