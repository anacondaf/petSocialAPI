pipeline {
    agent any

    environment {
        GREET_ENV = 'HELLO WORLD'
    }

    stages{
        stage("Build") {
            steps {
                echo "Say ${GREET_ENV}"
                sh 'printenv'
            }
        }
    }
}