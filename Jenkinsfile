pipeline {
    agent any

    environment {
      PATH = $PATH:/usr/local/go/bin  
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Check-Env') {
          steps {
            echo $PATH
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

