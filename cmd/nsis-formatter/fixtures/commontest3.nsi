
${For} $R1 1 5
                ${If} $R1 == 2
                                ${Select} $test10
                                                ${Case} 'a'
                                                                DetailPrint '$$test10 is equal to a'
                                                ${Case} 'b'
                                                                DetailPrint '$$test10 is equal to b'
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
                                                ${Case} 'c'
                                                                DetailPrint '$$test10 is equal to c'
                                                ${CaseElse}
                                                                DetailPrint '$$test10 is equal to $test10'
                                ${EndSelect}
                ${ElseIf} $R1 == 4
                                ${Select} $test10
                                                ${Case} 'a'
                                                                DetailPrint '$$test10 is equal to a'
                                                ${Case} 'b'
                                                                DetailPrint '$$test10 is equal to b'
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
                                                ${Case} 'c'
                                                                DetailPrint '$$test10 is equal to c'
                                                ${CaseElse}
                                                                DetailPrint '$$test10 is equal to $test10'
                                ${EndSelect}
                ${EndIf}
                DetailPrint $R1
${Next}
