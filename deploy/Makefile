.DEFAULT_GOAL:=help
SHELL:=/bin/bash
NAMESPACE=monitoring

##@ Application

install: ## Install all resources (CR/CRD's, RBAC and Operator)
	
	@echo ....... Applying Rules and Service Account .......
	- kubectl apply -f role.yaml -n ${NAMESPACE}
	- kubectl apply -f role_binding.yaml  -n ${NAMESPACE}
	- kubectl apply -f service_account.yaml  -n ${NAMESPACE}
	@echo ....... Applying Operator .......
	- kubectl apply -f deploy.yaml -n ${NAMESPACE}
	@echo ....... service .......
	- kubectl apply -f service.yaml -n ${NAMESPACE}
	@echo ....... servicemonitor .......
	- kubectl apply -f servicemonitor.yaml -n ${NAMESPACE}
	@echo ....... prometheus .......
	- kubectl apply -f prometheus.yaml -n ${NAMESPACE}

uninstall: ## Install all resources (CR/CRD's, RBAC and Operator)
	
	@echo ....... Cleaning .......
	@echo ....... Rules and Service Account .......
	- kubectl delete -f role.yaml -n ${NAMESPACE}
	- kubectl delete -f role_binding.yaml  -n ${NAMESPACE}
	- kubectl delete -f service_account.yaml  -n ${NAMESPACE}
	@echo ....... Operator .......
	- kubectl delete -f deploy.yaml -n ${NAMESPACE}
	@echo ....... service .......
	- kubectl delete -f service.yaml -n ${NAMESPACE}
	@echo ....... servicemonitor .......
	- kubectl delete -f servicemonitor.yaml -n ${NAMESPACE}
	@echo ....... prometheus .......
	- kubectl delete -f prometheus.yaml -n ${NAMESPACE}

