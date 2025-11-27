pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "",
            artifactNumToKeepStr: "10",
            daysToKeepStr: "",
            numToKeepStr: "10"
        )
    }
    triggers {
        cron "H 8 * * 3"
    }
    stages {
        stage("Clear") {
            steps {
                dir("/data/automation/checkouts"){
                    script {
                        deleteDir()
                    }
                }
            }
        }
        stage("Checkout"){
            steps{
                dir("/data/automation/checkouts/platypus"){
                    git url: "https://github.com/farghul/platypus.git", branch: "main"
                }
            }
        }
        stage("Build") {
            steps {
                dir("/data/automation/checkouts/platypus"){
                    script {
                        sh "go build -o /data/automation/bin/platypus"
                    }
                }
            }
        }
        stage("Run") {
            steps {
                dir("/data/automation/checkouts/platypus"){
                    script {
                        sh "./platypus.sh"
                    }
                }
            }
        }
    }
}