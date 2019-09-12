TEMPLATE_NAME := test
TEST_PROJECT_DIR := ../my-project4

clean:
	@rm -rf ${TEST_PROJECT_DIR}

save:
	boilr template save -f . ${TEMPLATE_NAME}

test: clean save
	boilr template use ${TEMPLATE_NAME} ${TEST_PROJECT_DIR}