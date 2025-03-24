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
        stage("Pull Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/platypus") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git pull
                        '''
                    }
                }
            }
        }
        stage("Build Platypus") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/platypus") {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/platypus ."
                    }
                }
            }
        }
        stage("Run Platypus") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            sh "/data/automation/scripts/platypus.sh"
                        }
                    }
                }
            }
        }
    }
}