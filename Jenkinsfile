pipeline {
    agent any
    tools { go '1.22.5' }
    options {
        parallelsAlwaysFailFast()
    }
    stages {
        stage('checkout') {
            steps {
                git branch: 'master', url: 'https://github.com/maradamark99/todo-go.git'
            }
        }
        stage('build') {
            steps {
                sh 'go build -o ./out/app *.go'
            }
        }
        stage('test') {
            parallel {
                stage('unit tests') {
                    steps {
                        sh 'go test'
                    }
                }
                stage('coverage') {
                    steps {
                        sh 'go test -coverprofile ./out/cover.out'   
                    }
                }
            }
        }
        stage('archive') {
            steps {
                archiveArtifacts artifacts: '**/app,**/*.out'
            }
        }
    }
    post {
        changed { /* runs everytime current outcome is diff from prev */
            emailext body: "${env.BUILD_URL}\n${currentBuild.absoluteUrl}", 
            recipientProviders: [previous()], 
            subject: "${currentBuild.currentResult}: Job ${env.JOB_NAME} #${env.BUILD_NUMBER}", 
            to: "test@test.com"    
        }
    }
}