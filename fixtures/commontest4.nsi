!define NETVersion "4.6.2"
!define NETInstaller "ndp462-web.exe"
Section "MS .NET Framework v${NETVersion}" SecFramework
                Var /GLOBAL dotNET462IsThere
                ReadRegDWORD $dotNET462IsThere HKLM "SOFTWARE\Microsoft\NET Framework Setup\NDP\v4\Full" "Release"
                IntCmp $dotNET462IsThere 394802 is_equal is_less is_greater

              is_equal:
                Goto check_net_framework

              is_greater:
                Goto check_net_framework

              is_less:
                Goto net_framework_install

              net_framework_install:
                DetailPrint "Starting Microsoft .NET Framework v${NETVersion} Setup..."
                ExecWait "$INSTDIR\${NETInstaller}"
                Goto check_net_framework

              check_net_framework:
                ExecWait '"$INSTDIR\LetsPRO.exe" check'
                StrCpy $0 0
                Goto waitting_check_finish

              waitting_check_finish:
                nsProcess::_FindProcess "LetsPRO.exe"
                Pop $R0
                IntCmp $R0 0 check_in_progress check_finish check_finish

              check_in_progress:
                IntOp $0 $0 + 1
                ${While} $0 > 12
                                Goto check_finish
                ${EndWhile}
                Sleep 300
                Goto waitting_check_finish

              check_finish:
                StrCpy $0 0
                GoTo process_check_result

              process_check_result:
                MessageBox MB_OK "check finish"
                Goto done

              done:
SectionEnd
