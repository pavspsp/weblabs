@echo off

for /l %%i in (1,1,5) do (
	echo "Client start"
	start "Python" cmd /c "python client.py %%i & timeout /t 10 /nobreak"
	pause
)
goto end


:exit
echo enddd
exit /b 0


:end
echo end
eXit /b 0