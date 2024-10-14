import jdk.internal.agent.resources.agent
pipeline {
    agent {
        docker {
            image 'golang:1.22-alpine'
        }
    }
    environment {
        GO114MODULE ='on'
    }
    stages {
        stage('Build') {
            steps {
                sh 'pwd'
                sh 'go mod tidy'
                sh 'go build -o bin/docker-ui main.go'
            }
        }
        stage('Test') {
            steps {
                sh 'pwd'
                sh 'ls bin'
            }
        }
    }
}
