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
        stage("Empty_Folder") {
            steps {
                dir("/data/automation/checkouts"){
                    script {
                        deleteDir()
                    }
                }
            }
        }
        stage("Checkout_Platypus"){
            steps{
                dir("/data/automation/checkouts/platypus"){
                    git url: "https://github.com/farghul/platypus.git" , branch: "main"
                }
            }
        }
        stage("Build_Platypus") {
            steps {
                dir("/data/automation/checkouts/platypus"){
                    script {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/platypus"
                    }
                }
            }
        }
        stage("Checkout_DAC") {
            steps{
                dir("/data/automation/checkouts/dac"){
                    git credentialsId: "DES-Project", url: "https://bitbucket.org/bc-gov/desso-automation-conf.git", branch: "main"
                }
            }
        }
        stage("Run_Platypus") {
            steps {
                dir("/data/automation/checkouts/dac/scripts/plugin"){
                    script {
                        sh "./platypus.sh"
                    }
                }
            }
        }
    }
}