#!groovy

podTemplate(label: "myPod", containers: [containerTemplate(name: "golang", image: "golang:1.8.0-alpine", ttyEnabled: true, command: "cat")]) {
    node("myPod") {
        container("golang") {
            stage("Checkout") {
                checkout scm
            }

            stage("Run Go Test") {
                sh "go test -v"
            }
        }
    }
}
