!define DEBUG_PATH "E:\NSIS_Test\TmpProgram"
!define OUTPUT_PATH "E:\NSIS_Test\Output"
!define INSTALL_PATH "E:\NSIS_Test\Output"

!include LogicLib.nsh

Name "NSIS_VariableConstants_Test"
Caption "NSIS_VariableConstants_Test"

InstallDir ${INSTALL_PATH}

OutFile "Galatea.exe"
Section "My Program"
                SetOutPath ${OUTPUT_PATH}
                File /r "${DEBUG_PATH}\*.*"

                ; ----------------- If - ElseIf - Else - EndIf -----------------

                var /GLOBAL test1
                StrCpy $test1 'a'

                ${If} $test1 == 'a'
                                DetailPrint '$$test1 is equal to a'
                ${ElseIf} $test1 == 'b'
                                DetailPrint '$$test1 is equal to b'
                ${Else}
                                DetailPrint '$$test1 is equal to $test1'
                ${EndIf}

                ; ----------------- IfNot - Else - EndIf -----------------

                var /GLOBAL test2
                StrCpy $test2 'z'

                ${IfNot} $test2 == 'a'
                                DetailPrint '$$test2 is not equal to a'
                ${Else}
                                DetailPrint '$$test2 is equal to $test2'
                ${EndIf}

                ; ----------------- If - ElseIf - ElseIfNot - EndIf -----------------

                var /GLOBAL test3
                StrCpy $test3 'b'

                ${If} $test3 == 'a'
                                DetailPrint '$$test3 is equal to a'
                ${ElseIfNot} $test3 == 'b'
                                DetailPrint '$$test3 is not equal to b'
                ${Else}
                                DetailPrint '$$test3 is equal to $test3'
                ${EndIf}

                ; ----------------- IfNot & Unless -----------------
                var /GLOBAL test4
                StrCpy $test4 'z'

                ${Unless} $test4 == 'a'
                                DetailPrint '$$test4 is not equal to a'
                ${Else}
                                DetailPrint '$$test4 is equal to $test4'
                ${EndUnless}

                ; ----------------- ElseIfNot & ElseUnless -----------------

                var /GLOBAL test5
                StrCpy $test5 'b'

                ${If} $test5 == 'a'
                                DetailPrint '$$test5 is equal to a'
                ${ElseUnless} $test5 == 'b'
                                DetailPrint '$$test5 is not equal to b'
                ${Else}
                                DetailPrint '$$test5 is equal to $test5'
                ${EndIf}

                ; ----------------- AndIf & AndIfNot & AndUnless -----------------

                var /GLOBAL test6
                StrCpy $test6 'a'

                ${If} $test6 == 'a'
                                DetailPrint '$$test6 is equal to a'
                ${AndIf} $test6 != 'b'
                                DetailPrint '$$test6 is not equal to b'
                ${AndIfNot} $test6 == 'b'
                                DetailPrint '$$test6 is not equal to b'
                ${AndUnless} $test6 == 'b'
                                DetailPrint '$$test6 is not equal to b'
                ${EndIf}

                ; ----------------- OrIf & OrIfNot & OrUnless -----------------

                var /GLOBAL test7
                StrCpy $test7 'z'

                ${If} $test7 == 'a'
                                DetailPrint '$$test7 is equal to a'
                ${OrIf} $test7 == 'b'
                                DetailPrint '$$test7 is equal to b'
                ${OrIfNot} $test7 != 'b'
                                DetailPrint '$$test7 is equal to b'
                ${OrUnless} $test7 != 'b'
                                DetailPrint '$$test7 is equal to b'
                ${Else}
                                DetailPrint '$$test7 is equal to $test7'
                ${EndIf}

                ; ----------------- IfThen -----------------

                var /GLOBAL test8
                StrCpy $test8 'a'
                ${IfThen} $test8 == 'a' ${|} Goto x1 ${|}

              x1:
                DetailPrint 'x1'
                Goto endOfxy1

              y1:
                DetailPrint 'y1'
                Goto endOfxy1

              endOfxy1:
                DetailPrint 'endOfxy1'

                ; ----------------- IfNotThen -----------------
                var /GLOBAL test9
                StrCpy $test9 'b'
                ${IfNotThen} $test9 == 'a' ${|} Goto y2 ${|}

              x2:
                DetailPrint 'x2'
                Goto endOfxy2

              y2:
                DetailPrint 'y2'
                Goto endOfxy2

              endOfxy2:
                DetailPrint 'endOfxy2'

                ; ----------------- IfCmd -----------------

                ${IfCmd} MessageBox MB_YESNO "MY_YESNO" /SD IDYES IDYES ${||} Goto y3 ${|}

              x3:
                DetailPrint 'x3'
                Goto endOfxy3

              y3:
                DetailPrint 'y3'
                Goto endOfxy3

              endOfxy3:
                DetailPrint 'endOfxy3'

                ; ----------------- Select - Case - CaseElse - EndSelect -----------------

                var /GLOBAL test10
                StrCpy $test10 'b'

                ${Select} $test10
                                ${Case} 'a'
                                                DetailPrint '$$test10 is equal to a'
                                ${Case} 'b'
                                                DetailPrint '$$test10 is equal to b'
                                ${Case} 'c'
                                                DetailPrint '$$test10 is equal to c'
                                ${CaseElse}
                                                DetailPrint '$$test10 is equal to $test10'
                ${EndSelect}

                ; ----------------- Select - Case - Default - EndSelect -----------------

                var /GLOBAL test11
                StrCpy $test11 'd'

                ${Select} $test11
                                ${Case} 'a'
                                                DetailPrint '$$test11 is equal to a'
                                ${Case} 'b'
                                                DetailPrint '$$test11 is equal to b'
                                ${Case} 'c'
                                                DetailPrint '$$test11 is equal to c'
                                ${Default}
                                                DetailPrint '$$test11 is equal to $test11'
                ${EndSelect}

                ; ----------------- Switch - Case - CaseElse - EndSwitch -----------------

                var /GLOBAL test12
                StrCpy $test12 'b'

                ${Switch} $test12
                                ${Case} 'a'
                                                DetailPrint '$$test12 is equal to a'
                                                ${Break}
                                ${Case} 'b'
                                                DetailPrint '$$test12 is equal to b'
                                                ${Break}
                                ${Case} 'c'
                                                DetailPrint '$$test12 is equal to c'
                                                ${Break}
                                ${CaseElse}
                                                DetailPrint '$$test12 is equal to $test12'
                ${EndSwitch}

                ; ----------------- Switch - Case - Default - EndSwitch -----------------
                var /GLOBAL test13
                StrCpy $test13 'b'

                ${Switch} $test13
                                ${Case} 'a'
                                                DetailPrint '$$test13 is equal to a'
                                                ${Break}
                                ${Case} 'b'
                                                DetailPrint '$$test13 is equal to b'
                                                ${Break}
                                ${Case} 'c'
                                                DetailPrint '$$test13 is equal to c'
                                                ${Break}
                                ${Default}
                                                DetailPrint '$$test13 is equal to $test13'
                ${EndSwitch}

                ; ----------------- Switch - Case - CaseElse - EndSwitch WithOut Break -----------------

                var /GLOBAL test14
                StrCpy $test14 'a'

                ${Switch} $test14
                                ${Case} 'a'
                                                DetailPrint '$$test14 is equal to a'
                                ${Case} 'b'
                                                DetailPrint '$$test14 is equal to b'
                                ${Case} 'c'
                                                DetailPrint '$$test14 is equal to c'
                                ${Default}
                                                DetailPrint '$$test14 is equal to $test14'
                ${EndSwitch}

                ; ----------------- While - EndWhile -----------------

                StrCpy $R1 0
                ${While} $R1 < 5
                                IntOp $R1 $R1 + 1
                                DetailPrint $R1
                ${EndWhile}

                ; ----------------- DoWhile - Loop -----------------

                StrCpy $R1 0
                ${DoWhile} $R1 < 5
                                IntOp $R1 $R1 + 1
                                DetailPrint $R1
                ${Loop}

                ; ----------------- DoUntil - Loop -----------------

                StrCpy $R1 0
                ${DoUntil} $R1 >= 5
                                IntOp $R1 $R1 + 1
                                DetailPrint $R1
                ${Loop}

                ; ----------------- Do - LoopWhile -----------------

                StrCpy $R1 0
                ${Do}
                                IntOp $R1 $R1 + 1
                                DetailPrint $R1
                ${LoopWhile} $R1 < 5

                ; ----------------- Do - LoopUntil -----------------

                StrCpy $R1 0
                ${Do}
                                IntOp $R1 $R1 + 1
                                DetailPrint $R1
                ${LoopUntil} $R1 >= 5

                ; ----------------- Break & Continue -----------------

                StrCpy $R1 0
                ${While} $R1 < 5
                                IntOp $R1 $R1 + 1
                                ${If} $R1 == 2
                                                ${Continue}
                                ${ElseIf} $R1 == 4
                                                ${Break}
                                ${EndIf}
                                DetailPrint $R1
                ${EndWhile}

                ; ----------------- ExitDo -----------------

                StrCpy $R1 0
                ${Do}
                                IntOp $R1 $R1 + 1
                                ${If} $R1 == 4
                                                ${ExitDo}
                                ${EndIf}
                                DetailPrint $R1
                ${LoopWhile} $R1 < 5

                ; ----------------- ExitWhile -----------------

                StrCpy $R1 0
                ${While} $R1 < 5
                                IntOp $R1 $R1 + 1
                                ${If} $R1 == 4
                                                ${ExitWhile}
                                ${EndIf}
                                DetailPrint $R1
                ${EndWhile}

                ; ----------------- For - Next -----------------

                ${For} $R1 1 5
                                DetailPrint $R1
                ${Next}

                ; ----------------- ForEach -----------------

                ${ForEach} $R1 1 5 + 1
                                DetailPrint $R1
                ${Next}
                ${ForEach} $R1 10 2 - 2
                                DetailPrint $R1
                ${Next}

                ; ----------------- Break & Continue -----------------

                ${For} $R1 1 5
                                ${If} $R1 == 2
                                                ${Continue}
                                ${ElseIf} $R1 == 4
                                                ${Break}
                                ${EndIf}
                                DetailPrint $R1
                ${Next}

                ; ----------------- ExitFor -----------------

                ${For} $R1 1 5
                                ${If} $R1 == 4
                                                ${ExitFor}
                                ${EndIf}
                                DetailPrint $R1
                ${Next}

SectionEnd
