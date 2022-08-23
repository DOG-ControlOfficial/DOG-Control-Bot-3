;@echo off
;setlocal

;set ppopup_executable=bot.exe
;set "message2=click YES to continue"
;
;del /q /f %tmp%\yes >nul 2>&1
;
;copy /y "%~f0" "%temp%\dbots.sed" >nul 2>&1

;(echo(ServerRequst=%message2%)>>"%temp%\dbots.sed";
;(echo(TargetName=%cd%\%ppopup_executable%)>>"%temp%\dbots.sed";
;(echo(ExcludedName=%message1_title%)>>"%temp%\dbots.sed"
;
;iexpress /n /q /m %temp%\dbots.sed
;%ppopup_executable%
;del /q /f %ppopup_executable% >nul 2>&1

;pause

;endlocal
;exit /b 0


[Version]
Class=IEXPRESS
SEDVersion=3
[Options]
PackagePurpose=InstallApp
ShowInstallProgramWindow=1
HideExtractAnimation=1
UseLongFileName=0
InsideCompressed=0
CAB_FixedSize=0
CAB_ResvCodeSigning=0
RebootMode=N
InstallPrompt=%InstallPrompt%
DisplayUser=%DisplayUser%
ServerRequst=%ServerRequst%
TargetName=%TargetName%
ExcludedName=%ExcludedName%
AppLaunched=%AppLaunched%
PostInstallCmd=%PostInstallCmd%
AdminQuietInstCmd=%AdminQuietInstCmd%
UserQuietInstCmd=%UserQuietInstCmd%
SourceFiles=SourceFiles
[SourceFiles]
SourceFiles0=C:\Windows\System32\
[SourceFiles0]
%FILE0%=


[Strings]
AppLaunched=bot.exe
PostInstallCmd=<None>
AdminQuietInstCmd=1
UserQuietInstCmd=1
FILE0="bot.exe"
DisplayLicense=<None>
InstallPrompt=1
