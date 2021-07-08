document.addEventListener('load', () => GraphQLPlayground.init(document.getElementById('root'), {
    endpoint: 'http://localhost:6654/graphql',
    'general.betaUpdates': true,
    'tracing.hideTracingResponse': false,
    'tracing.tracingSupported': false
}));
