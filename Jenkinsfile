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
        stage("Pull Config Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/bitbucket/desso-automation-conf") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git switch main
                        git pull
                        '''
                    }
                }
            }
        }
        stage("Pull Program Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/platypus") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git switch main
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
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/platypus"
                    }
                }
            }
        }
        stage("Run Platypus") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            dir("/data/automation/bitbucket/desso-automation-conf/scripts/plugin") {
                                sh "./platypus.sh"
                            }
                        }
                    }
                }
            }
        }
    }
}