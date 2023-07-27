pipeline {
    agent any

    environment {
        GREET = 'HELLO WORLD'
    }

    stages{
        stage("Build") {
            steps {
                echo "Say ${GREET}"
                sh 'printenv'
            }
        }
    }
}