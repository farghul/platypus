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
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/platypus"
                    }
                }
            }
        }
        stage("Run") {
            steps {
                dir("/data/automation/checkouts/platypus"){
                    script {
                        sh "/data/automation/bin/platypus -r"
                    }
                }
            }
        }
    }
}