select name, akas
from azure.azure_network_security_group
where name = '{{resourceName}}' and resource_group = '{{resourceName}}'
