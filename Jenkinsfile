pipeline {
    agent {
        docker {
            image 'golang:1.17' // replace with your Go version
            args '-v /root/.cache:/root/.cache' // cache go build to speed up
        }
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
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

