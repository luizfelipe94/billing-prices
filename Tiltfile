docker_build('billing.localhost/billing-prices', '.')

k8s_yaml('k8s/deployment.yaml')

k8s_resource('billing-prices', port_forwards=8082)