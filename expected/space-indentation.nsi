# Scaffolding by https://github.com/idleberg/vscode-nsis

# Settings ---------------------------------
Name "demo"
OutFile "demo.exe"
RequestExecutionLevel user
Unicode true
InstallDir "$PROGRAMFILES\Demo"

# Includes ---------------------------------
!include "LogicLib.nsh"

# Pages ------------------------------------
PageEx license
  LicenseText "Readme"
PageExEnd
Page components
Page directory
Page instfiles

# Macros
!macro Macro
  Nop
!macroend

# Sections ---------------------------------
SectionGroup "LogicLib.nsh"
  Section "If"
    ${If} "true" == "true"
      Nop
    ${EndIf}

    ${If} "true" == "true"
    ${AndIf} "false" == "false"
      Nop
    ${EndIf}

    ${If} "true" == "true"
    ${OrIf} "false" == "false"
      Nop
    ${EndIf}

    ${If} "true" == "true"
      Nop
    ${ElseIf} "false" == "false"
    ${AndIfNot} "true" == "false"
      Nop
    ${ElseIfNot} "true" == "false"
    ${OrIfNot} "false" == "true"
      Nop
      Nop
    ${ElseUnless} "true" != "false"
      Nop
    ${Else}
      Nop
    ${EndIf}
  SectionEnd

  Section "For"
    ${For} $0 1 5
      Nop
    ${Next}
  SectionEnd

  Section "Switch"
    ${Switch} $R1
      ${Case} 1
        Nop
        ${Break}
      ${Case} 2
        Nop
        ${Break}
      ${Default}
        Nop
        ${Break}
    ${EndSwitch}
  SectionEnd
SectionGroupEnd

# Functions --------------------------------
Function .onInit
  !if "true" == "true"
    MessageBox MB_OK "Condition met"
  !endif

  !if "true" == "true"
    MessageBox MB_OK "Condition met"
  !else
    MessageBox MB_OK "Condition notmet"
  !endif

  !if "false" == "false"
    !if "true" != "false"
      MessageBox MB_OK "Condition met"
    !endif
  !endif
FunctionEnd
