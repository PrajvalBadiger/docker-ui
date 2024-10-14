import jdk.internal.agent.resources.agent
pipeline {
    agent {
        docker {
            image 'golang:1.22-alpine'
        }
    }
    parameters {
        string defaultValue: 'main', description: 'specify which branch to use.', name: 'BRANCH', trim: true
    }
    environment {
        GO114MODULE ='on'
        GOCACHE="${WORKSPACE}"
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
