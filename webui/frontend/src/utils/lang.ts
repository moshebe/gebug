export default {
    placeholders: {
        name: 'awesome-app',
        outputBinary: '/app',
        buildCommand: 'go build -o {{.output_binary}}',
        runCommand: '{{.output_binary}}',
        runtimeImage: 'golang:latest',
        debuggerPort: 4321,
        debuggerEnabled: false,
        exposePorts: 'PORT[:PORT]',
        networks: 'private-network',
        envName: 'FOO',
        envValue: 'BAR',  
    },
    labels: {
        name: 'Name',
        outputBinary: 'Output Binary Path',
        buildCommand: 'Build Command',
        runCommand: 'Run Command',        
        runtimeImage: 'Runtime Image',
        debuggerPort: 'Debugger Port',
        debuggerEnabled: 'Debugger Enabled',
        exposePorts: 'Expose Ports',
        networks: 'Networks',
        environment: 'Environment',
    },
    addLabels: {
        networks: '+ Add Network',
        environment: '+ Add Environment Variable',
        exposePorts: '+ Add Port',
    },
    help: {
        debuggerPort: 'Delve debugger port',
    }
};