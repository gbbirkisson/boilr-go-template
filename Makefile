TEMPLATE_NAME := test

save:
	boilr template save -f . ${TEMPLATE_NAME}

test: save
	@rm -rf ./my-project
	boilr template use ${TEMPLATE_NAME} ./my-project