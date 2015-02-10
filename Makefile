
default: parser
	@echo ""
	@echo -e "\033[1;37mThank you for using the 'Greatest Compiler in America'"
	@echo -e "                               \033[1;36m~~ Bill Clinton\033[0;00m"
	@echo -e ""
	@echo -e "\033[0;31mAMERICA \033[1;37mAMERICA \033[0;34mAMERICA \033[0;00m"
	@echo -e "\033[1;37m_____________--____----"
	@echo -e "|\033[1;34m * * * *\033[1;37m |\033[1;37m__\--\__\--\033[1;37m|"
	@echo -e "|\033[1;34m* * * * *\033[1;37m|\033[0;31m___\--\__\-\033[1;37m|"
	@echo -e "|\033[1;34m_*_*_*_*_\033[1;37m|\033[1;37m\___\--\___\\033[1;37m|"
	@echo -e "|\033[0;31m___________\___\--\__\033[1;37m|"
	@echo -e "|\033[1;37m____________\___\--\_\033[1;37m|"
	@echo -e "|\033[0;31m_____________\___\---\\033[1;37m|"
	@echo -e "||"
	@echo -e "||"
	@echo -e "||"
	@echo -e "||"
	@echo -e "||"
	@echo -e "||"
	@echo -e "||"
	@echo -e "||"
	@echo -e ""

parser: y.tab.c lex.yy.c
	gcc y.tab.c lex.yy.c -o parser -lfl

y.tab.c: yacc.y
	bison -y -d -g -t --verbose yacc.y

lex.yy.c: lex.l
	lex lex.l

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output parser.out parser

