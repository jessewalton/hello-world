pipeline {
    agent any

    environment {
     	GOPATH = "/home/vagrant/go"
      PATH = "${env.PATH}:/usr/local/go/bin:/home/vagrant/go/bin"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Check-Env') {
          steps {
            echo "${env.PATH}"
          }
        }
        stage('Lint') {
            steps {
                sh 'go get -u golang.org/x/lint/golint'
                sh 'golint ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o hello-world'
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'hello-world', fingerprint: true
        }
    }
}

