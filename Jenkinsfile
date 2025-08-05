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
        stage("Checkout Platypus") {
            steps {
                checkout([$class: 'GitSCM',
                    branches: [[name: "main"]],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [[$class: 'RelativeTargetDirectory',
                    relativeTargetDir: 'platypus']],
                    submoduleCfg: [],
                    userRemoteConfigs: [[url: 'https://github.com/farghul/platypus.git']]
                ])
            }
        }
        stage("Build Platypus") {
            steps {
                script {
                    sh "/data/apps/go/bin/go build -o /data/automation/bin/platypus platypus/."
                }
            }
        }
        stage("Checkout DAC") {
            steps {
                checkout([$class: 'GitSCM',
                    branches: [[name: "main"]],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [[$class: 'RelativeTargetDirectory',
                    relativeTargetDir: 'desso']],
                    submoduleCfg: [],
                    userRemoteConfigs: [[credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git']]
                ])
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