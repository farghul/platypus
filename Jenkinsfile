pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "28",
            artifactNumToKeepStr: "5",
            daysToKeepStr: "56",
            numToKeepStr: "10"
        )
    }
    triggers {
        cron "H 8 * * 3"
    }
    stages {
        stage('Clean WS') {
            steps {
                cleanWs()
            }
        }
        stage('Checkout Platypus'){
            steps{
                dir('/data/automation/temp/platypus'){
                    git url: 'https://github.com/farghul/platypus.git' , branch: 'main'
                }
            }
        }
        stage('Build Platypus') {
            steps {
                script {
                    sh "/data/apps/go/bin/go build -o /data/automation/bin/platypus"
                }
            }
        }
        stage("Checkout DAC") {
            steps{
                dir('/data/automation/temp/dac'){
                    git credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git', branch: 'main'
                }
            }
        }
        stage('Run Platypus') {
            steps {
                script {
                    sh './scripts/plugin/platypus.sh'
                }
            }
        }
    }
}