package frames

func init() {
	FrameMap["cuby"] = DefaultFrameType(_cuby)
}

var _cuby = []string{
	`        ...',,;;::cclooddoollcc:;;,,'....
       cxxxxxxxxxxxxxkk0Oddddoooolllllllc
       oxxxxxxxxxxxxxOXMMXkllllllllllllll
       dxxxxxxxxxxxONMMMMMMNkllllllllllll
       xxxxxxxxxxONMMMMMMMMMMNkllllllllll.
      .xxxxxxxx0NMMMMMMMMMMMMMMNOolllllll.
      .xxxxxx0WMMMMMMMMMMMMMMMMMMWOolllll'
      ,xxxx0WMMMMMMMMMMMMMMMMMMMMMMW0olll;
      :xxKWMMMMMMMMMMMMMMMMMMMMMMMMMMW0ol:
      lKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0o
      cKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXo
      .;oKMMMMMMMMMMMMMMMMMMMMMMMMMMMMKl'.
       ;;;l0WMMMMMMMMMMMMMMMMMMMMMMM0c...
       ;;;;;cOWMMMMMMMMMMMMMMMMMMWO:.....
       ';;;;;;ckWMMMMMMMMMMMMMMWk;.......
       .;;;;;;;;:xNMMMMMMMMMMNx,.........
        ..',;;;;;;:dXMMMMMMNd'.........
               ...',;oKMMKo'....
                       '`,

	`        ..'',,;;::cclloodoollcc:;;,,'....
       cxxxxxxxxxxxxxkk0xddddoooooolllllc
       dxxxxxxxxxxxxkKWMW0ollllllllllllll
       xxxxxxxxxxxkXMMMMMMW0ollllllllllll
      .xxxxxxxxxkXMMMMMMMMMMW0ollllllllll
      ,xxxxxxxkXMMMMMMMMMMMMMMW0dllllllll
      cxxxxxOXMMMMMMMMMMMMMMMMMMMKdllllll.
      oxxxONMMMMMMMMMMMMMMMMMMMMMMMKxllll.
      xxONMMMMMMMMMMMMMMMMMMMMMMMMMMMXxll.
     .ONMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXx.
     .kWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWO.
      ,:xNMMMMMMMMMMMMMMMMMMMMMMMMMMMWk;'
      .;;:xXMMMMMMMMMMMMMMMMMMMMMMMNd,'''
       ;;;;;dXMMMMMMMMMMMMMMMMMMMXo'.'''.
       ';;;;;;oKMMMMMMMMMMMMMMMXo'.'.''..
       .;;;;;;;;o0MMMMMMMMMMM0c'.''.'.''.
        ...',;;;;;l0WMMMMMWO:..'.''.....
               ...',cOWMNx;.....
                      .`,

	`        ..'',,;;::ccllooddollcc:;;,''...
       cddddxxxxxxxxxk0kddddddoooooollllc
       ddddddddddddx0NMWKxllllllllllllllc
      .ddddddddddx0WMMMMMMXxllllllllllllc
      ,ddddddddx0WMMMMMMMMMMXxllllllllllc
      cddddddxKWMMMMMMMMMMMMMMXkllllllllc
      dddddxKWMMMMMMMMMMMMMMMMMMNkllllllc
     .dddkKWMMMMMMMMMMMMMMMMMMMMMMNkolllc
     ;dkXMMMMMMMMMMMMMMMMMMMMMMMMMMMNOolc
     oXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNkc
     cKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXc
      ;l0MMMMMMMMMMMMMMMMMMMMMMMMMMMMKl'.
      ';;lOWMMMMMMMMMMMMMMMMMMMMMMM0c'''.
      .;;;;cOWMMMMMMMMMMMMMMMMMMWk;'''''.
       ,;;;;;:kWMMMMMMMMMMMMMMWk;'''''''.
       .;;;;;;;:kNMMMMMMMMMMNd,'''''''''.
         ..',;;;;:xNMMMMMMXo''''''''....
               ...';dXMWOl'.....
                      ,`,

	`        ..'',,;;::ccllloodoolcc:;;,''...
       :dddddxxxxxxxxOOxdddddddoooooolllc
       ddddddddddddkXMMXkollllllllllllllc
      .ddddddddddONMMMMMMNOollllllllllllc
      :ddddddddOXMMMMMMMMMMNOollllllllllc
      dddddddONMMMMMMMMMMMMMMNOollllllllc
     .dddddONMMMMMMMMMMMMMMMMMMWOollllllc
     :ddx0NMMMMMMMMMMMMMMMMMMMMMMW0ollllc
     dx0WMMMMMMMMMMMMMMMMMMMMMMMMMMW0oll:
    .0WMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0o:
     xNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWk,
     .;dNMMMMMMMMMMMMMMMMMMMMMMMMMMMWk;''
      ,;;dXMMMMMMMMMMMMMMMMMMMMMMMXd,'''.
      .;;;;oXMMMMMMMMMMMMMMMMMMMKo''''''.
       ,;;;;;oXMMMMMMMMMMMMMMMKl''''''''.
       .;;;;;;;oKMMMMMMMMMMWOc''''''''''.
         ...,;;;;l0MMMMMMWk;'''''''''...
               ...'cOWMKd,'.....
                     .`,

	`        ..',,;;;::cclllooodollc:;;,''...
       ;ddddddxxxxxxk0xxddddddddooooooool
       dddddddddddxKWMWOdoooooooooooooool
      'dddddddddxKWMMMMMWKdoooooooooooooc
      ldddddddxKWMMMMMMMMMW0doooooooooooc
     .ddddddkKWMMMMMMMMMMMMMWKdooooooooo:.
     ;ddddkXMMMMMMMMMMMMMMMMMMMKdooooooo:.
     oddOXMMMMMMMMMMMMMMMMMMMMMMMKxooooo;.
    .dONMMMMMMMMMMMMMMMMMMMMMMMMMMMKxool;.
    oNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMKxl;.
    ,OWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXc;
     'cOWMMMMMMMMMMMMMMMMMMMMMMMMMMMKl',;
      ;;ckWMMMMMMMMMMMMMMMMMMMMMMWOc'''';
      .;;;:kWMMMMMMMMMMMMMMMMMMWk;'''''',
       ,;;;;:kWMMMMMMMMMMMMMMNx,''''''''.
        ;;;;;;:kNMMMMMMMMMMXo,''''''''''.
          ..';;;:xNMMMMMM0l''''''''''...
               ...,xNMNx:'......
                     ,`,

	`        .'',,;;;::cccllooodoolc::;,'....
       ,dddddddxxxxx0kxxdddddddddoooooooc
       oddddddddddONMMKxooooooooooooooooc
      'dddddddddONMMMMMMXkoooooooooooooo:.
      odddddddONMMMMMMMMMMXxoooooooooooo:.
     .dddddx0WMMMMMMMMMMMMMMXkoooooooool;.
     ldddd0WMMMMMMMMMMMMMMMMMMXkoooooool;,
    .ddx0WMMMMMMMMMMMMMMMMMMMMMMXkoooooc;;
    cxKWMMMMMMMMMMMMMMMMMMMMMMMMMMNkooo:;,
   .KWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXko:;'
    lKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWd;;.
     ,lKMMMMMMMMMMMMMMMMMMMMMMMMMMMNx;';;.
     .,;lKMMMMMMMMMMMMMMMMMMMMMMMKo,''';;
      .;;;lKMMMMMMMMMMMMMMMMMMM0l'''''',;
       ';;;;lKMMMMMMMMMMMMMMMOc'''''''','
        ,;;;;;lKMMMMMMMMMMWk:'''''''''''.
          ..',;;lKMMMMMMXd,'''''''''''..
               ...c0MM0l''.......
                    .`,

	`        .'',,;;:::ccclllooodolc::;,'....
       'ddddddddxxxO0xxxxddddddddddoooooc
       odddddddddkKWMNkoooooooooooooooooc
      ,ddddddddkXMMMMMMNOooooooooooooooo:.
      oddddddkXMMMMMMMMMMNOooooooooooool;'
     ,dddddkXMMMMMMMMMMMMMMNOooooooooooc;;
     odddkXMMMMMMMMMMMMMMMMMMWOoooooooo:;;
    ;ddONMMMMMMMMMMMMMMMMMMMMMMNOdooooo;;;.
    dONMMMMMMMMMMMMMMMMMMMMMMMMMMWOoool;;;.
   cNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOoc;;;
    dNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMO:;;,
     ;xNMMMMMMMMMMMMMMMMMMMMMMMMMMMOc,,;;.
     .,;xNMMMMMMMMMMMMMMMMMMMMMMWk:,,,,;;.
      .,,;xNMMMMMMMMMMMMMMMMMMNd,,,,,,,;;
       ',,,;xWMMMMMMMMMMMMMMXd,,,,,,,,,,,
        ',,,,;xNMMMMMMMMMM0l,,,,,,,,,,,,.
          ..',,;xWMMMMMWk:,,,,,,,,,,,''.
               ..,xNMXd;''.......
`,

	`        .'',,;;:::ccclllooodolcc:;,'...
       .ddddddddxxx0kxxxxxddddddddddoooo:
       lddddddddd0WMW0dooooooooooooooooo:
      ,dddddddd0WMMMMMW0dooooooooooooool;.
      odddddd0WMMMMMMMMMWKdooooooooooooc;,
     :ddddx0WMMMMMMMMMMMMMW0doooooooooo:;;
    .dddxKWMMMMMMMMMMMMMMMMMWKdoooooool;;;.
    ldxKWMMMMMMMMMMMMMMMMMMMMMW0doooooc;;;,
   'xKMMMMMMMMMMMMMMMMMMMMMMMMMMW0dooo:;;;'
   OWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0dl;;;;.
   .kWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXc;;;;
    .:OWMMMMMMMMMMMMMMMMMMMMMMMMMMXo,,;;;,
     .,:OWMMMMMMMMMMMMMMMMMMMMMM0l,,,,;;;.
      .,,:OMMMMMMMMMMMMMMMMMMWO:,,,,,,,;;
       .,,,c0MMMMMMMMMMMMMMWk:,,,,,,,,,;;
        .,,,,l0MMMMMMMMMMXd;,,,,,,,,,,,,.
           ..,,c0MMMMMM0l,,,,,,,,,,,,,''
               ..c0MWk:,''.......
                   .`,

	`        .'',,;;:::cccllllooooolc:;,'...
       .ddddddddxxOOxxxxxxxddddddddddddo;
       lddddddddkXMMXkoooooooooooooooooo;
      ,dddddddkNMMMMMMXxoooooooooooooool;.
      odddddkXMMMMMMMMMMXkooooooooooooo:;;
     cddddkXMMMMMMMMMMMMMMXkooooooooool;;;.
    'dddONMMMMMMMMMMMMMMMMMMXxooooooooc;;;,
    odONMMMMMMMMMMMMMMMMMMMMMMXxoooooo;;;;;
   cONMMMMMMMMMMMMMMMMMMMMMMMMMMXxoooc;;;;;
  .NMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMKdo:;;;;,
   ;0MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWo;;;;;.
    .lKMMMMMMMMMMMMMMMMMMMMMMMMMMWk:,;;;;;
     .,lKMMMMMMMMMMMMMMMMMMMMMMXd;,,,,;;;'
      .,,oXMMMMMMMMMMMMMMMMMMKo,,,,,,,;;;.
       .,,,dNMMMMMMMMMMMMMM0l,,,,,,,,,,;;
        .,,,;dNMMMMMMMMMWOc,,,,,,,,,,,,;.
           ..,;xNMMMMMXd;,,,,,,,,,,,,,,'
               .'dNMKo,,''........
`,

	`        .',,;;::::cccclllooooolc:;,'...
        ooooddddxk0xxxxxxxxxdddddddddddd,
       cooooooox0WMWOdddddooddoodooddodl;
      ,ooodoodKWMMMMMNOddddoodooddoodoo:;'
     .oooood0WMMMMMMMMMNOdddooodoodoodo;;;
     looox0WMMMMMMMMMMMMMNOdooddooddooc;;;'
    ;oodKWMMMMMMMMMMMMMMMMMNOdodooodoo;;;;;
   .oxKWMMMMMMMMMMMMMMMMMMMMMNkdoooodc;;;;;'
   dKMMMMMMMMMMMMMMMMMMMMMMMMMMNkdodo;;;;;;.
  cWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXxdc;;;;;;
   cXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMk;;;;;;'
    .oNMMMMMMMMMMMMMMMMMMMMMMMMMM0c,;;;;;;.
     .,dNMMMMMMMMMMMMMMMMMMMMMWk:,,,,;;;;,
      .,;xWMMMMMMMMMMMMMMMMMNx;,,,,,,,;;;.
       .,,;kWMMMMMMMMMMMMMXd;,,,,,,,,,;;;
         ',,:OWMMMMMMMMW0o,,,,,,,,,,,,,;.
           ..'c0MMMMMWk:,,,,,,,,,,,,,,,,
               .:0MNx;,,'''.......
`,

	`         ',,;;::::cccclllloooolc:;,'..
        loooooddx0Oxxxxxxxxxxxdddddddddd.
       :oooooooONMMKxddddddddddddddddddl;
      'ooooooONMMMMMWKxdddddddddddddddo:;'
     .oooookNMMMMMMMMMW0ddddddddddddddc;;;
     loookNMMMMMMMMMMMMMW0dddddddddddo;;;;,
    :ookNMMMMMMMMMMMMMMMMMW0dddddddddc;;;;;.
   ,okNMMMMMMMMMMMMMMMMMMMMMW0ddddddo;;;;;;;
  .kNMMMMMMMMMMMMMMMMMMMMMMMMMWOdddd:;;;;;;;
  OMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNkdo;;;;;;;.
   oNMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0:;;;;;;;
    'xWMMMMMMMMMMMMMMMMMMMMMMMMMXo;,;;;;;;.
     .;kWMMMMMMMMMMMMMMMMMMMMM0l,,,,;;;;;;
      .':OMMMMMMMMMMMMMMMMMWOc,,,,,,,;;;;.
        ',cKMMMMMMMMMMMMMWk:,,,,,,,,,,;;;
         .',oXMMMMMMMMMXd:,,,,,,,,,,,,;;.
           ..,dNMMMMMKl,,,,,,,,,,,,,,,,;
               .dWWOc,,,'''........
`,

	`         ',,;;::::cccclllllooooc:;,'..
        :oooooddk0xxxxxxxxxxxxxddddddddo.
       ;ooooooxKWMNOdddddddddddddddddddc;
      'ooooodKMMMMMMNkddddddddddddddddl::,
     .ooood0WMMMMMMMMMXxdddddddddddddd::::.
     lood0WMMMMMMMMMMMMWKxdddddddddddl::::;
    cod0WMMMMMMMMMMMMMMMMMKxddddddddo::::::'
   :d0WMMMMMMMMMMMMMMMMMMMMWKxddddddc:::::::.
  ;0WMMMMMMMMMMMMMMMMMMMMMMMMW0ddddl::::::::.
  XMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOdd::::::::,
  .dWMMMMMMMMMMMMMMMMMMMMMMMMMMMMXl::::::::
    'OMMMMMMMMMMMMMMMMMMMMMMMMMNx:,:::::::'
     .c0MMMMMMMMMMMMMMMMMMMMMXd;,,,;::::::
       'lXMMMMMMMMMMMMMMMMMKo,,,,,,,;::::'
        .,dNMMMMMMMMMMMMM0l,,,,,,,,,,;:::
         .';kWMMMMMMMMNkc,,,,,,,,,,,,,;:.
            .;OMMMMMNx;,,,,,,,,,,,,,,,,;
               ;0MKd;,,''''........
`,

	`         ',,;:::::ccccllllloooolc;,'..
        ,ooooodd0kxxxxxxxxxxxxxxxxdddddl
       'ooooooONMWKxdddddddddddddddddddc;
      .oooooONMMMMMW0dddddddddddddddddl::,
     .ooookNMMMMMMMMMNOddddddddddddddo::::.
    .lookXMMMMMMMMMMMMMXkddddddddddddc:::::
    lokXMMMMMMMMMMMMMMMMMNkdddddddddl::::::;
   ckXMMMMMMMMMMMMMMMMMMMMMXkddddddo::::::::'
  oXMMMMMMMMMMMMMMMMMMMMMMMMMKxddddc::::::::,
 .WMMMMMMMMMMMMMMMMMMMMMMMMMMMW0ddl:::::::::
  .kWMMMMMMMMMMMMMMMMMMMMMMMMMMMNo:::::::::.
    ,KMMMMMMMMMMMMMMMMMMMMMMMMWOc;::::::::;
     .lXMMMMMMMMMMMMMMMMMMMMNk:;,;;:::::::.
       'dWMMMMMMMMMMMMMMMMNx;;,;,;,;:::::,
        .;OMMMMMMMMMMMMMXd;;,;,;,;,;;::::
         ..c0MMMMMMMMW0l;,;,;,;,;,;,;;::.
            .lXMMMMWOc,;;,;,;,;,;,;,;,;;
              .lNNk:,,,,'''........
`,

	`         .,,;::::cccccclllllooolc;,'..
        .ooooodk0xxxxxxxxxxxxxxxxxxxxxdc
       .looooxKMMXkdddddddddddddddddddoc;
      .looodKMMMMMMXxdddddddddddddddddccc,
     .lood0WMMMMMMMMW0xddddddddddddddlcccc.
    .lod0WMMMMMMMMMMMMWOddddddddddddocccccc.
    loOWMMMMMMMMMMMMMMMMNOdddddddddoccccccc:
   lOWMMMMMMMMMMMMMMMMMMMMNOdddddddccccccccc;
  kWMMMMMMMMMMMMMMMMMMMMMMMMXxddddlccccccccc:
 ;WMMMMMMMMMMMMMMMMMMMMMMMMMMWKxdocccccccccc.
  .OMMMMMMMMMMMMMMMMMMMMMMMMMMMWxcccccccccc;
    ;XMMMMMMMMMMMMMMMMMMMMMMMMKl;:ccccccccc
     .oNMMMMMMMMMMMMMMMMMMMWOl;;;;:ccccccc.
       'kMMMMMMMMMMMMMMMMWO:;;;;;;;:ccccc,
        .cKMMMMMMMMMMMMNx:;;;;;;;;;;:cccc
          .oNMMMMMMMMXd;;;;;;;;;;;;;;:cc.
            .xWMMMMKo;;;;;;;;;;;;;;;;;:;
              'OW0l;,,,''''..........
`,

	`         .,,;::::cccccclllllloooc:,'.
        .llllodKkxxxxxxxxxxxxxxxxxxxxxd;
       .lllloOWMW0xdddddddddddddddddddol,
      .lllokNMMMMMNOdddddddddddddddddolll'
     .lllxNMMMMMMMMMXkdddddddddddddddlllll.
    .llxXMMMMMMMMMMMMWKxddddddddddddllllllc.
   .ldKMMMMMMMMMMMMMMMMW0xdddddddddllllllllc.
  .oKMMMMMMMMMMMMMMMMMMMMW0xddddddolllllllllc
 .0WMMMMMMMMMMMMMMMMMMMMMMMNkddddolllllllllll.
 lWMMMMMMMMMMMMMMMMMMMMMMMMMMKxddlllllllllll;
  .0MMMMMMMMMMMMMMMMMMMMMMMMMMWklllllllllll:
    :NMMMMMMMMMMMMMMMMMMMMMMMXd;:llllllllll.
     .xWMMMMMMMMMMMMMMMMMMMKo;;;;:clllllll.
       'KMMMMMMMMMMMMMMMM0l;;;;;;;;clllll;
         oNMMMMMMMMMMMWOc;;;;;;;;;;;clllc
          .kWMMMMMMMNk:;;;;;;;;;;;;;;cll.
            ,KMMMMNx:;;;;;;;;;;;;;;;;;c,
              :NXd;,,,,''''...........
`,

	`          ',;:::ccccccclllllllool:,..
         cllloO0kxxxxxxxxxxxxxxxxxxxxxo'
        clllxXMMXkxxxxxxxxxxxxxxxxxxxdol'
       clldKMMMMMMKxxxxxxxxxxxxxxxxxdolll'
      clo0WMMMMMMMMWOxxxxxxxxxxxxxxxolllll.
    .coOWMMMMMMMMMMMMXkxxxxxxxxxxxxolllllll.
   .lkNMMMMMMMMMMMMMMMMKxxxxxxxxxxdlllllllll.
  .xXMMMMMMMMMMMMMMMMMMMWKxxxxxxxdlllllllllll.
 ,XMMMMMMMMMMMMMMMMMMMMMMMWOxxxxdllllllllllll;
 oMMMMMMMMMMMMMMMMMMMMMMMMMMXkxdolllllllllllc
  .0MMMMMMMMMMMMMMMMMMMMMMMMMWOollllllllllll.
    cWMMMMMMMMMMMMMMMMMMMMMMNx::lllllllllll.
     .kMMMMMMMMMMMMMMMMMMMXd:;;;;lllllllll,
       ,XMMMMMMMMMMMMMMMKd;;;;;;;;cllllll:
         dMMMMMMMMMMMMKl;;;;;;;;;;;cllllc
          '0MMMMMMMM0l;;;;;;;;;;;;;;:lll.
            cNMMMW0c;;;;;;;;;;;;;;;;;:l.
             .xWOc;,,,,''''...........
`,

	`          .,::::cccccccclllllllllc,..
         ;llldKkkkkkkkkxxxxxxxxxxxxxxxl.
        ;lloOWMW0xxxxxxxxxxxxxxxxxxxxdol.
       :llkNMMMMMNkxxxxxxxxxxxxxxxxxdoooo.
      cldXMMMMMMMMMKxxxxxxxxxxxxxxxdoooooo.
    .cdKMMMMMMMMMMMMW0xxxxxxxxxxxxdoooooooo.
   .l0WMMMMMMMMMMMMMMMXkxxxxxxxxxxoooooooooo.
  .OWMMMMMMMMMMMMMMMMMMWKkxxxxxxxoooooooooooo'
 cNMMMMMMMMMMMMMMMMMMMMMMW0xxxxxoooooooooooool
 dMMMMMMMMMMMMMMMMMMMMMMMMMNkxxoooooooooooool.
  'KMMMMMMMMMMMMMMMMMMMMMMMMMKoooooooooooooo.
    lWMMMMMMMMMMMMMMMMMMMMMWk::looooooooooo'
     .OMMMMMMMMMMMMMMMMMMNk:;;;;cooooooooo;
       :NMMMMMMMMMMMMMMNx:;;;;;;;cooooooo:
        .OMMMMMMMMMMMNd;;;;;;;;;;;:looool
          ;NMMMMMMMXd;;;;;;;;;;;;;;:lool.
            xMMMMXo;;;;;;;;;;;;;;;;;;co.
             'KKo;,,,,,'''''..........
`,

	`          .,:::ccccccccclllllllllc;'.
         .lllk0kkkkkkkkkkkxxxxxxxxxxxxc
        ,lldKMMXkxxxxxxxxxxxxxxxxxxxxdol.
       ;lo0MMMMMW0xxxxxxxxxxxxxxxxxxdoooo.
      :lkWMMMMMMMMNkxxxxxxxxxxxxxxxdoooooo.
    .cxNMMMMMMMMMMMMKxxxxxxxxxxxxxdoooooooo.
   .oKMMMMMMMMMMMMMMMWOxxxxxxxxxxdoooooooooo'
  ,0WMMMMMMMMMMMMMMMMMMXkxxxxxxxdoooooooooooo,
 oWMMMMMMMMMMMMMMMMMMMMMMKxxxxxdoooooooooooooo.
 dMMMMMMMMMMMMMMMMMMMMMMMMWOxxdoooooooooooooo'
  'KMMMMMMMMMMMMMMMMMMMMMMMMXdoooooooooooooo,
    oMMMMMMMMMMMMMMMMMMMMMW0c:loooooooooooo;
     'KMMMMMMMMMMMMMMMMMWOc;;;;coooooooooo:
       oWMMMMMMMMMMMMMWk:;;;;;;;:ooooooooc
        'XMMMMMMMMMMWk:;;;;;;;;;;;loooool
          oWMMMMMMWk:;;;;;;;;;;;;;;coool
           .KMMMWk:;;;;;;;;;;;;;;;;;:oo
             lXx:,,,,,,''''''.........
`,

	`           ,:cccccccccccclllllllll;'.
         .cloKkkkkkkkkkkkkkkkkkkkxkxxd;
        .clkWMW0xxxxxxxxxxxxxxxxxxxxxddc
       ,lxNMMMMMXkxxxxxxxxxxxxxxxxxxddddl
      ;oKMMMMMMMMW0xxxxxxxxxxxxxxxxddddddo.
     cOWMMMMMMMMMMMNkxxxxxxxxxxxxxddddddddo.
   .xNMMMMMMMMMMMMMMWKxxxxxxxxxxxddddddddddd,
  ;KMMMMMMMMMMMMMMMMMMNOxxxxxxxxddddddddddddd:
 xMMMMMMMMMMMMMMMMMMMMMMXkxxxxxddddddddddddddd'
 dMMMMMMMMMMMMMMMMMMMMMMMW0xxxddddddddddddddd:
  'XMMMMMMMMMMMMMMMMMMMMMMMXxddddddddddddddd:
    dMMMMMMMMMMMMMMMMMMMMMKl:lddddddddddddd:
     ,XMMMMMMMMMMMMMMMMW0o;;;;cddddddddddd:
       xMMMMMMMMMMMMMWOc;;;;;;;:oddddddddc
        :NMMMMMMMMMM0l;;;;;;;;;;;cddddddc
         .OMMMMMMM0l;;;;;;;;;;;;;;:odddc
           :WMMM0l;;;;;;;;;;;;;;;;;;ldl
            .OOc,,,,,,,'''''''.......
`,

	`           ':cccccccccccclllllllll:'.
          ;cx0kkkkkkkkkkkkkkkkkkkkkkko.
        .:oKMMXkxxxxxxxxxxxxxxxxxxxxxdd;
       .lOWMMMMW0xxxxxxxxxxxxxxxxxxxddddc
      ,dNMMMMMMMMXkxxxxxxxxxxxxxxxxddddddl
     cKMMMMMMMMMMMW0kxxxxxxxxxxxxxxdddddddd.
   .kWMMMMMMMMMMMMMMXkxxxxxxxxxxxdxddddddddd,
  :XMMMMMMMMMMMMMMMMMW0kxxxxxxxxddddddddddddd:
 kMMMMMMMMMMMMMMMMMMMMMXkkxxxxxddddddddddddddd:
 oMMMMMMMMMMMMMMMMMMMMMMWKkxxxdddddddddddddddl
  'XMMMMMMMMMMMMMMMMMMMMMMNxxdddddddddddddddc
    xMMMMMMMMMMMMMMMMMMMMXo;lddddddddddddddc
     ;NMMMMMMMMMMMMMMMMKo;;;;codddddddddddc
      .0MMMMMMMMMMMMMKl;;;;;;;;ldddddddddc
        oMMMMMMMMMMXo;;;;;;;;;;;:dddddddc
         'XMMMMMMXd;;;;;;;;;;;;;;;lddddc
           xMMMNd;;;;;;;;;;;;;;;;;;:od:
            ;Kd;,,,,,,'''''''''......
`,

	`           .:cccccccccccccclllllllc'
          'l0Okkkkkkkkkkkkkkkkkkkkkkkl
         ;xNMW0kkkkkkkkkkkkkkxkxkxxxxxd.
       .lKMMMMMXkkkkkkkkkkkkkkkkxkxxxxxx;
      'kWMMMMMMMWOkkkkkkkkkkkkkkkxxxxxxxxc
     lXMMMMMMMMMMMXkkkkkkkkkkkkkxxxxxxxxxxo.
   .OMMMMMMMMMMMMMMWOkkkkkkkkkkxxxxxxxxxxxxx,
  cNMMMMMMMMMMMMMMMMMKkkkkkkkkxxxxxxxxxxxxxxxc
 OMMMMMMMMMMMMMMMMMMMMNOkkkkkxxxxxxxxxxxxxxxxxl
 lWMMMMMMMMMMMMMMMMMMMMWKkkkxxxxxxxxxxxxxxxxxo
  'XMMMMMMMMMMMMMMMMMMMMMNkxxxxxxxxxxxxxxxxxo
    kMMMMMMMMMMMMMMMMMMMXd:lxxxxxxxxxxxxxxxl
     cWMMMMMMMMMMMMMMMXd;;;;:oxxxxxxxxxxxxc
      .KMMMMMMMMMMMMNd:;;;;;;;cdxxxxxxxxxc
        kMMMMMMMMMNx:;;;;;;;;;;:oxxxxxxx:
         :WMMMMMWk:;;;;;;;;;;;;;;cdxxxx;
          .KMMWOc;;;;;;;;;;;;;;;;;:lxx,
            dk:,,,,,,,''''''''''''..'
`,

	`            :cccccccccccccccclllllc'
          .o0kkkkkkkkkkkkkkkkkkkkkkkx;
         ,0WWXkkkkkkkkkkkkkkkkkkkkkkxxl.
        oNWWWWW0kkkkkkkkkkkkkkkkkkkxxxxx'
      .0WWWWWWWWKkkkkkkkkkkkkkkkkkxxxxxxx:
     lNWWWWWWWWWWNOkkkkkkkkkkkkkkkxxxxxxxxo.
   .0WWWWWWWWWWWWWW0kkkkkkkkkkkkxxxxxxxxxxxx'
  lNWWWWWWWWWWWWWWWWXOkkkkkkkkkxxxxxxxxxxxxxxc
 0WWWWWWWWWWWWWWWWWWWW0kkkkkkxxxxxxxxxxxxxxxxxo
 :WWWWWWWWWWWWWWWWWWWWWKkkkkxxxxxxxxxxxxxxxxxx.
  .KWWWWWWWWWWWWWWWWWWWWNkkxxxxxxxxxxxkxxxxxd.
    kWWWWWWWWWWWWWWWWWWXd:lxxxxxxxxxxxxxxxxo
     cWWWWWWWWWWWWWWWNx:;;;:oxxxxxxxxxxxxxc
      'XWWWWWWWWWWWWk:;;;;;;;cdxxxxxxxxxx:
       .0WWWWWWWWWOc;;;;;;;;;;:lxxxxxxxx;
         oWWWWWW0l;;;;;;;;;;;;;;:oxxxxx'
          ;NWWKo;;;;;;;;;;;;;;;;;;cdxx.
           .kl,,,,,,'''''''''''''''';
`,

	`            ;cccccccccccccccccccccc;
           xOkkkkkkkkkkkkkkkkkkkkkkkd.
         ,XWW0kkkkkkkkkkkkkkkkkkkkkkkk:
        dWWWWWXkkkkkkkkkkkkkkkkkkkkkkkkd.
      .KWWWWWWWNOkkkkkkkkkkkkkkkkkkkkkkkx,
     oWWWWWWWWWWW0kkkkkkkkkkkkkkkkkkkkkkkkl
   .KWWWWWWWWWWWWWXkkkkkkkkkkkkkkkkkkkkkkkkx'
  lWWWWWWWWWWWWWWWWNOkkkkkkkkkkkkkkkkkkkkkkkkc
 0WWWWWWWWWWWWWWWWWWWKkkkkkkkkkkkkkkkkkkkkkkkkd
 ,NWWWWWWWWWWWWWWWWWWWXOkkkkkkkkkkkkkkkkkkkkkk'
  .KWWWWWWWWWWWWWWWWWWWNkkkkkkkkkkkkkkkkkkkkx.
    kWWWWWWWWWWWWWWWWWNd:lxkkkkkkkkkkkkkkkko
     lWWWWWWWWWWWWWWNk:;;;:okkkkkkkkkkkkkkc
      ,NWWWWWWWWWWWOc;;;;;;;:dkkkkkkkkkkk;
       .XWWWWWWWWKl;;;;;;;;;;;cxkkkkkkkk,
         kWWWWWXd:;;;;;;;;;;;;;:lxkkkkx.
          oWWNx:;;;;;;;;;;;;;;;;;:okkd.
           ;x;''''''''''''''''''''';
`,

	`            'cccccccccccccccccccccc:
          .kkkkkkkkkkkkkkkkkkkkkkkkko.
         :NWXOkkkkkkkkkkkkkkkkkkkkkkkx,
       .kWWWWN0kkkkkkkkkkkkkkkkkkkkkkkkl
      ,XWWWWWWWKkkkkkkkkkkkkkkkkkkkkkkkkx'
     dWWWWWWWWWWXOkkkkkkkkkkkkkkkkkkkkkkkkc
   .KWWWWWWWWWWWWNOkkkkkkkkkkkkkkkkkkkkkkkkx.
  lWWWWWWWWWWWWWWWN0kkkkkkkkkkkkkkkkkkkkkkkkkc
 OWWWWWWWWWWWWWWWWWWKOkkkkkkkkkkkkkkkkkkkkkkkkx
 .XWWWWWWWWWWWWWWWWWWXOkkkkkkkkkkkkkkkkkkkkkkk;
  .0WWWWWWWWWWWWWWWWWWNOkkkkkkkkkkkkkkkkkkkkx.
    kWWWWWWWWWWWWWWWWNx:lxkkkkkkkkkkkkkkkkkd
     oWWWWWWWWWWWWWWOc::::lkkkkkkkkkkkkkkkc
      ;NWWWWWWWWWWKl::::::::okkkkkkkkkkkk;
       'NWWWWWWWNd:::::::::::cdkkkkkkkkk.
        .KWWWWNkc::::::::::::::cxkkkkkd.
          OWW0c::::::::::::::::::lxkkl
           lc'''''''''''''''',,,,,,c
`,

	`            .clllllccccccccccccccccc
          :kkkkkkkkkkkkkkkkkkkkkkkkOc.
         xWWKkkkkkkkkkkkkkkkkkkkkkOOOd'
       .KWWWWXOkkkkkkkkkkkkkkkkkkOOOOOkc
      cNWWWWWWXOkkkkkkkkkkkkkkkOOOOOOOOOd.
    .kWWWWWWWWWN0kkkkkkkkkkkkkOOOOOOOOOOOk:
   'XWWWWWWWWWWWN0kOkkkkkkkkOOOOOOOOOOOOOOOx.
  oNWWWWWWWWWWWWWWKOOOOkkkkOOOOOOOOOOOOOOOOOO:
 kWWWWWWWWWWWWWWWWWXOOOOOOOOOOOOOOOOOOOOOOOOOOx.
 .KWWWWWWWWWWWWWWWWWNOOOOOOOOOOOOOOOOOOOOOOOOO:
   OWWWWWWWWWWWWWWWWWNOOOOOOOOOOOOOOOOOOOOOOk'
    xWWWWWWWWWWWWWWWNk:lxOOOOOOOOOOOOOOOOOOd.
     oWWWWWWWWWWWWW0l::::lxOOOOOOOOOOOOOOOc
      cWWWWWWWWWWXd::::::::lkOOOOOOOOOOOk'
       ;NWWWWWWNk::::::::::::okOOOOOOOOx.
        'XWWWWKl:::::::::::::::dOOOOOOl
         .KWNd::::::::::::::::::cdOOO;
          .o,''''''''''''',,,,,,,,:o
`,

	`            .:llllllllcccccccccccccc.
          dkkkkOOOOOOOOOOOOOOkkkkkkk,.
        .0NNOkkkkkkkkkkkkkkkkkkkkOOOOl'
       :NNNNN0OOOkkkkkkkkkkkkkkkOOOOOOx;.
      dNNNNNNN0OOOOOOOOOOkkkkkOOOOOOOOOOo.
    .ONNNNNNNNNKOOOOOOOOOOOOOOOOOOOOOOOOOk:
   ,XNNNNNNNNNNNXOOOOOOOOOOOOOOOOOOOOOOOOOOd.
  oNNNNNNNNNNNNNNXOOOOOOOOOOOOOOOOOOOOOOOOOOO:
 oNNNNNNNNNNNNNNNNNOOOOOOOOOOOOOOOOOOOOOOOOOOOx.
  ONNNNNNNNNNNNNNNNN0OOOOOOOOOOOOOOOOOOOOOOOOOl
   kNNNNNNNNNNNNNNNNN0OOOOOOOOOOOOOOOOOOOOOOk'
    dNNNNNNNNNNNNNNNO:lxOOOOOOOOOOOOOOOOOOOd.
     oNNNNNNNNNNNNKo::::lxOOOOOOOOOOOOOOOO:
      lNNNNNNNNNNx::::::::lxOOOOOOOOOOOOk.
       cNNNNNNNKl:::::::::::lkOOOOOOOOOd
        :NNNNXx:::::::::::::::okOOOOOO:
         ,NNOc::::::::::::::::::oOOOk.
          ':''''''''''',,,,,,,,,,;l
`,

	`            .;lllllllllcccccccccccc:.
         .xkOOOOOOOOOOOOOOOOOOOOOOOo'.
        ;XNKOOOOOOOOOOOOOOOOOkkkOO00x:'.
       oNNNNXOOOOOOOOOOOOOOOOOOO000000o,.
      kNNNNNNXOOOOOOOOOOOOOOOOO00000000Oc.
    .0NNNNNNNNXOOOOOOOOOOOOOOO00000000000x;
   ,XNNNNNNNNNNXOOOOOOOOOOOO000000000000000o.
  lNNNNNNNNNNNNNN0OOOOOOOO000000000000000000k;
 :NNNNNNNNNNNNNNNN0OOOOOO000000000000000000000d.
  dNNNNNNNNNNNNNNNN0OOO00000000000000000000000l
   dNNNNNNNNNNNNNNNN0O0000000000000000000000O,
    oNNNNNNNNNNNNNNOccx00000000000000000000d.
     oNNNNNNNNNNNXd::::lxO000000000000000O:
      lNNNNNNNNNOc:::::::cx0000000000000x.
       oNNNNNNXd:::::::::::cx0000000000l
        lNNNNOc::::::::::::::lx000000O,
         lNXd::::::::::::::::::lk000d
          ;'...'''''''',,,,,,,,,;cx
`,

	`            .;cllllllllcccccccccccc:'
         :kOOOOOOOOOOOOOOOOOOOOOOOO:''.
        oNN0OOOOOOOOOOOOOOOOOOOOO000o,'.
       kNNNN0OOOOOOOOOOOOOOOOOOO00000Oc'.
     .ONNNNNN0OOOOOOOOOOOOOOOO000000000x;.
    .KNNNNNNNN0OOOOOOOOOOOOO0000000000000o,
   ,XNNNNNNNNNN0OOOOOOOOOOO000000000000000Ol.
  cNNNNNNNNNNNNNKOOOOOOOO0000000000000000000k,
 .NNNNNNNNNNNNNNNKOOOOOO0000000000000000000000o
  :NNNNNNNNNNNNNNNKOOO000000000000000000000000o
   cNNNNNNNNNNNNNNN0O00000000000000000000000O,
    lNNNNNNNNNNNNN0ccx000000000000000000000d
     lNNNNNNNNNNXx::::cx00000000000000000O;
      oNNNNNNNNKl:::::::cd00000000000000d.
       dNNNNNNk:::::::::::cd00000000000:
        dNNNKo::::::::::::::cd0000000x.
         xNOc:::::::::::::::::cd0000c
          ;.....''''''',,,,,,;;;:dx
`,

	`           ..,clllllllllcccccccccc::'.
         dkOOOOOOOOOOOOOOOOOOOOOOOx,,,.
        kXKOOOOOOOOOOOOOOOOOOOOO00Kkc,,.
      .0XXXXOOOOOOOOOOOOOOOOOOO0KKKKKx;,'
     .0XXXXXKOOOOOOOOOOOOOOOO0KKKKKKKK0o,'
    .KXXXXXXXKOOOOOOOOOOOOOO0KKKKKKKKKKKOl,.
   ,XXXXXXXXXXKOOOOOOOOOOO0KKKKKKKKKKKKKKKk:.
  :XXXXXXXXXXXXKOOOOOOOO00KKKKKKKKKKKKKKKKK0x'
  KXXXXXXXXXXXXXKOOOOOO0KKKKKKKKKKKKKKKKKKKKK0l
  'XXXXXXXXXXXXXXKOOO0KKKKKKKKKKKKKKKKKKKKKKKKd
   ,XXXXXXXXXXXXXXKO000KKKKKKKKKKKKKKKKKKKKKO,
    :XXXXXXXXXXXXKlcd00KKKKKKKKKKKKKKKKKKKKo
     cXXXXXXXXXXk::::cd0KKKKKKKKKKKKKKKKKO,
      oXXXXXXXXd::::::::dOKKKKKKKKKKKKKKo
       xXXXXXKl:::::::::::oOKKKKKKKKKKO,
        kXXXk:::::::::::::::oOKKKKKKKo
         OXo::::::::::::::::::oOKKKO'
         .'......''''',,,,,,;;;;oO
`,

	`           ..,:lllllllllccccccccc::;'.
        .xkOOOOOOOOOOOOOOOOOOOOOOOl,,,'
       .0X0OOOOOOOOOOOOOOOOOOOOO0K0d;,,'
      .KXXX0OOOOOOOOOOOOOOOOOO0KKKKK0o,,'
     .KXXXXX0OOOOOOOOOOOOOOO0KKKKKKKKKOc,,
    .KXXXXXXX0OOOOOOOOOOOOO0KKKKKKKKKKKKk:,.
   'KXXXXXXXXXOOOOOOOOOOO0KKKKKKKKKKKKKKKKx;.
  'KXXXXXXXXXXXOOOOOOOO0KKKKKKKKKKKKKKKKKKK0d.
  kXXXXXXXXXXXXXOOOOOOKKKKKKKKKKKKKKKKKKKKKKK0c
   0XXXXXXXXXXXXXOOO0KKKKKKKKKKKKKKKKKKKKKKKKKd
   .KXXXXXXXXXXXXKOKKKKKKKKKKKKKKKKKKKKKKKKKO,
    'XXXXXXXXXXXKo:d0KKKKKKKKKKKKKKKKKKKKKKl
     :XXXXXXXXX0c::::dOKKKKKKKKKKKKKKKKKKk.
      oXXXXXXXk::::::::oOKKKKKKKKKKKKKK0c
       kXXXXXd:::::::::::oOKKKKKKKKKKKk.
        OXXKl::::::::::::::lkKKKKKKK0:
        .KOc:::::::::::::::::lkKKKKd.
         .........'''',,,,,;;;;lkO
`,

	`           ..,:coolllllllccccccc:::;,.
        ,kkkOOOOOOOOOOOOOOOOOOOOOx:,,,,
       ,KXOkkkkkkkkkkOOOOOOOOOO0KKOc,,,,
      'KXXKOkkOOOOOOOOOOOOOOO0KKKKKKk:,,,
     .KXXXXKOOOOOOOOOOOOOOOOKKKKKKKKKKx;,,.
    .KXXXXXXKOOOOOOOOOOOOO0KKKKKKKKKKKKKd;,.
   .KXXXXXXXX0OOOOOOOOOO0KKKKKKKKKKKKKKKK0o,.
  .0XXXXXXXXXX0OOOOOOO0KKKKKKKKKKKKKKKKKKKK0l.
  :XXXXXXXXXXXX0OOOOOKKKKKKKKKKKKKKKKKKKKKKKKO;
   xXXXXXXXXXXXXOOO0KKKKKKKKKKKKKKKKKKKKKKKKKKd
    0XXXXXXXXXXXKOKKKKKKKKKKKKKKKKKKKKKKKKKKO'
    .KXXXXXXXXXKd:oOKKKKKKKKKKKKKKKKKKKKKKKc
     ;XXXXXXXXKl::::oOKKKKKKKKKKKKKKKKKKKx.
      lXXXXXX0c:::::::lkKKKKKKKKKKKKKKK0;
       kXXXXOc::::::::::lkKKKKKKKKKKKKd.
       .0XXk::::::::::::::cxKKKKKKKKO,
        .Kd:::::::::::::::::cxKKKKKc
         .........'''',,,,;;;;:d0d
`,

	`           ..,;cooollllllcccccc::::;,..
        lkkkkkkOOkOkkOOOOOOOOOOOOl:,,,,.
       :KKOkkkkkkkkkkkkkkkkkkk0KXKd;,,,,.
      ;KKK0kkkkkkkkkkkkkkkkOOKXXXXXKd;,,,.
     .KKKKK0kkkkkkkkOOOOOOO0KXXXXXXXX0o,,,.
    .0KKKKKKOkkOOOOOOOOOO0KXXXXXXXXXXXXOl,,.
   .0KKKKKKKKOOOOOOOOOO0KXXXXXXXXXXXXXXXXOc,.
   kKKKKKKKKKKOOOOOOOOKXXXXXXXXXXXXXXXXXXXKkc.
  .KKKKKKKKKKK0OOOOOKXXXXXXXXXXXXXXXXXXXXXXXKk.
   :KKKKKKKKKKK0OO0KXXXXXXXXXXXXXXXXXXXXXXXXXXo
    xKKKKKKKKKKKkKXXXXXXXXXXXXXXXXXXXXXXXXXXO'
     0KKKKKKKKKx:lOKXXXXXXXXXXXXXXXXXXXXXXKc
     'KKKKKKKKd::::lkKXXXXXXXXXXXXXXXXXXKd.
      lKKKKKKo:::::::ckKXXXXXXXXXXXXXXXO,
       kKKKKo::::::::::cxKXXXXXXXXXXXKo.
       .KK0l::::::::::::::d0XXXXXXXKx.
        ;0c:::::::::::::::::o0XXXXO,
         .........'''',,,;;;::oOK
`,

	`           ..,;clooolllllcccccc::::;,..
        dkkkkkkkkkkkkkkkkkkOOOOOxc:;;;;.
       oK0kkkkkkkkkkkkkkkkkkkOKXXOl;;;;;.
      :KKKOkkkkkkkkkkkkkkkkO0XXXXXXOc;;;;.
     .KKKKKOkkkkkkkkkkkkkk0KXXXXXXXXXk:;;,.
    .OKKKKK0kkkkkkkkkkkO0KXXXXXXXXXXXXKx:;,.
    kKKKKKKK0kkkkkkkkkOKXXXXXXXXXXXXXXXXKx:,.
   oKKKKKKKKKOkkkkkOOKXXXXXXXXXXXXXXXXXXXXKx;
   kKKKKKKKKKKOOOOO0XXXXXXXXXXXXXXXXXXXXXXXXKd.
   .KKKKKKKKKK0OO0KXXXXXXXXXXXXXXXXXXXXXXXXXXXc
    cKKKKKKKKKKkKXXXXXXXXXXXXXXXXXXXXXXXXXXXk.
     kKKKKKKKKk:lkKXXXXXXXXXXXXXXXXXXXXXXX0:
     .KKKKKKKk::::cxKXXXXXXXXXXXXXXXXXXXKo.
      cKKKKKk:::::::cxKXXXXXXXXXXXXXXXXk'
       kKKKk:::::::::::d0XXXXXXXXXXXX0:.
       .KKk::::::::::::::oOXXXXXXXXKo.
        ck:::::::::::::::::lkXXXXXx.
         .........''',,,;;;;:lkKk'
`,

	`          ...,;:loooolllllccccc::::;,..
        xkkkkkkkkkkkkkkkkkkkkkOOoc:;;;;'
       dKOkkkkkkkkkkkkkkkkkkk0XXKd:;;;;;.
      :KK0kkkkkkkkkkkkkkkkk0KXXXXXKd;;;;;.
     .0KKKOkkkkkkkkkkkkkkOKXXXXXXXXXKo;;;;.
     OKKKKKOkkkkkkkkkkkOKXXXXXXXXXXXXX0o;;;.
    oKKKKKK0kkkkkkkkkOKXXXXXXXXXXXXXXXXX0o;;
   ;KKKKKKKK0kkkkkkO0XXXXXXXXXXXXXXXXXXXXX0o,
   :KKKKKKKKKOkkkO0XXXXXXXXXXXXXXXXXXXXXXXXX0l
    OKKKKKKKK0Ok0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX,
    'KKKKKKKKKOKXXXXXXXXXXXXXXXXXXXXXXXXXXXXx.
     dKKKKKKKOccxKXXXXXXXXXXXXXXXXXXXXXXXX0;
     .0KKKKKOc:::cd0XXXXXXXXXXXXXXXXXXXXKl.
      :KKKK0c:::::::o0XXXXXXXXXXXXXXXXXo.
       kKK0l::::::::::o0XXXXXXXXXXXXXk,.
       .K0l:::::::::::::lkKXXXXXXXXO:..
        oo::::::::::::::::cxKXXXXKl.
           ........'',,,;;;:cd0Kl
`,

	`          ...,;:cloooollllcccc::::;;,'.
       .kkkkkkkkkkkkkkkkkkkkkk0xlc:;;;;,
       x0kkkkkkkkkkkkkkkkkkkOKXXOl;;;;;;'
      :00OkkkkkkkkkkkkkkkkOKXXXXXXOl;;;;;.
     .0000kkkkkkkkkkkkkkO0XXXXXXXXXXkc;;;;.
     x00000kkkkkkkkkkkOKXXXXXXXXXXXXXXkc;;;.
    :000000Okkkkkkkkk0XXXXXXXXXXXXXXXXXXkc;;
   .00000000kkkkkkk0XXXXXXXXXXXXXXXXXXXXXXkc,
   .00000000Okkkk0XXXXXXXXXXXXXXXXXXXXXXXXXXk;
    l00000000OkOKXXXXXXXXXXXXXXXXXXXXXXXXXXXXK.
     O0000000O0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXd.
     :0000000l:dKXXXXXXXXXXXXXXXXXXXXXXXXXk,
      O00000o::::oOXXXXXXXXXXXXXXXXXXXXXO:.
      ,0000d:::::::lOXXXXXXXXXXXXXXXXXKc.
       k00x::::::::::lkXXXXXXXXXXXXXXo'.
       '0k:::::::::::::cxKXXXXXXXXXx,..
        dc::::::::::::::::d0XXXXXk;.
            .......'',,,;;::oOXO;
`,

	`          ..',;:cloooollllcccc::::;,,'.
       'kkkkkkkkkkkkkkkkkkkkkOOocc:;;;;;
       kOkkkkkkkkkkkkkkkkkkk0XNKd:;;;;;;,
      :00kkkkkkkkkkkkkkkkk0XNNNNNKd:;;;;;'
      O00Okkkkkkkkkkkkkk0XNNNNNNNNNKd:;;;;.
     l0000kkkkkkkkkkkk0XNNNNNNNNNNNNNKd:;;;
    .00000Okkkkkkkkk0XNNNNNNNNNNNNNNNNNKd:;,
    x000000Okkkkkk0XNNNNNNNNNNNNNNNNNNNNNKd:'
    d0000000kkkkOXNNNNNNNNNNNNNNNNNNNNNNNNNKd.
    .0000000OkOKNNNNNNNNNNNNNNNNNNNNNNNNNNNNNO
     x000000OONNNNNNNNNNNNNNNNNNNNNNNNNNNNNXo
     '000000o:o0NNNNNNNNNNNNNNNNNNNNNNNNNXx.
      x0000x::::lkXNNNNNNNNNNNNNNNNNNNNNk,.
      '000O:::::::cxXNNNNNNNNNNNNNNNNNO;..
       x00l:::::::::cxKNNNNNNNNNNNNN0c..
       '0d:::::::::::::d0NNNNNNNNNKl...
        d::::::::::::::::lOXNNNNXo..
            .......'',,;;;:ckXXd.
                             .`,

	`         ...',;:clooooollllccc::::;,,'..
       ;kkkkkkkkkkkkkkkkkkkkk0xlcc::::::.
       kkkkkkkkkkkkkkkkkkkk0XNNOl:::::::;
      ;OOkkkkkkkkkkkkkkkkOXNNNNNNOl::::::'
      kOOkkkkkkkkkkkkkkOXNNNNNNNNNNOc:::::.
     ;OOOOkkkkkkkkkkkOXNNNNNNNNNNNNNNOl:::;
     kOOOOkkkkkkkkkOXNNNNNNNNNNNNNNNNNNOl::,
    ;OOOOOOkkkkkkOKNNNNNNNNNNNNNNNNNNNNNNOl:.
    'OOOOOOkkkkOKNNNNNNNNNNNNNNNNNNNNNNNNNNOl
     kOOOOOOkkKNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNo
     :OOOOOOkXNNNNNNNNNNNNNNNNNNNNNNNNNNNNNKc
      OOOOOx:lkXNNNNNNNNNNNNNNNNNNNNNNNNNXo.
      lOOOOc:::cxXNNNNNNNNNNNNNNNNNNNNNXd'.
      .OOOo:::::::dKNNNNNNNNNNNNNNNNNNd'..
       xOk::::::::::d0NNNNNNNNNNNNNNk,...
       ,Oc::::::::::::oONNNNNNNNNNO;...
        l:::::::::::::::cxXNNNNN0:..
             ......'',,;;::dKN0:.
`,

	`         ...',;:ccloooollllccc:::;;,''..
       :kkkkkkkkkkkkkkkkkkkkOOollc::::::.
       kkkkkkkkkkkkkkkkkkkOKNNKd::::::::;
      ,OkkkkkkkkkkkkkkkkOKNNNNNNKd:::::::,
      dOOkkkkkkkkkkkkkOKNNNNNNNNNNKd::::::.
     .OOOkkkkkkkkkkkOKNNNNNNNNNNNNNNKd::::;
     lOOOOkkkkkkkkkKNNNNNNNNNNNNNNNNNNXd:::'
     OOOOOkkkkkkkKNNNNNNNNNNNNNNNNNNNNNNXxc:.
     kOOOOkkkkkKNNNNNNNNNNNNNNNNNNNNNNNNNNXx:
     cOOOOOkkKNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNX,
     .OOOOOxXNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN0;
      xOOOk::xXNNNNNNNNNNNNNNNNNNNNNNNNNNKc.
      ;OOOo::::d0NNNNNNNNNNNNNNNNNNNNNNKc..
       OOk:::::::o0NNNNNNNNNNNNNNNNNNKl...
       dOl:::::::::lONNNNNNNNNNNNNNXo....
       ,k::::::::::::ckXNNNNNNNNNXd'....
        c;::::::::::::::d0NNNNNNd'..
             ......'',,;::oONNd.
                            .`,

	`         ...',;;:clodooolllccc:::;;,''..
       cxxxkkkkkkkkkkkkkkkkk0dollcc:::::.
       kxxxxxxxxxxxxxxkkkk0NWNkc:::::::::
      'kkxxxxxxxxxxxxkkk0NWWWWWNOl:::::::,
      lkkxxxxxxxxxxkkk0NWWWWWWWWWNkc::::::.
      kkkxxxxxxxxkkk0NWWWWWWWWWWWWWNOl::::;
     'kkkkxxxxxxkk0NWWWWWWWWWWWWWWWWWNOl:::.
     lkkkkxxxxkk0NWWWWWWWWWWWWWWWWWWWWWNOo::
     :kkkkkxkk0NWWWWWWWWWWWWWWWWWWWWWWWWWW0o'
     .kkkkkk0NWWWWWWWWWWWWWWWWWWWWWWWWWWWWWN0
      xkkkx0WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWNk'
      ckkkl:o0WWWWWWWWWWWWWWWWWWWWWWWWWWNk;.
      .kkx::::lONWWWWWWWWWWWWWWWWWWWWWNk;..
       kko::::::ckNWWWWWWWWWWWWWWWWWWO;...
       lk:::::::::cxXWWWWWWWWWWWWWW0:....
       ,o::::::::::::dKWWWWWWWWWW0c.....
        ;;:::::::::::::oONWWWWW0c....
              .....'',;;:ckXW0:.
                           .`,

	`        ....',;;:clooooolllccc:::;;,''..
       lxxxxxxkkkkkkkkkkkkk0koollcc:::::'
       xxxxxxxxxxxxxxxxxxOXWW0d::::::::::
      .kxxxxxxxxxxxxxxx0NWWWWWWKd::::::::,
      ,kxxxxxxxxxxxxxONWWWWWWWWWWKd:::::::
      lkkxxxxxxxxxx0NWWWWWWWWWWWWWWKdc::::,
      kkkxxxxxxxx0NWWWWWWWWWWWWWWWWWWXxc:::
     .kkkxxxxxk0NWWWWWWWWWWWWWWWWWWWWWWXxc:,
      kkkkxxx0NWWWWWWWWWWWWWWWWWWWWWWWWWWXxc.
      dkkkxONWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWXo
      ckkkkWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWNd.
      'kkd:cONWWWWWWWWWWWWWWWWWWWWWWWWWWNd'
       kkl:::cxXWWWWWWWWWWWWWWWWWWWWWWNd'..
       dk:::::::dKWWWWWWWWWWWWWWWWWWNd'...
       cd:::::::::oKWWWWWWWWWWWWWWNx'....
       'l:::::::::::lONWWWWWWWWWNx,.....
        ,;;:::::::::::ckNWWWWWNx'....
              .....',,;;:dKWNd'.
                           '`,

	`        ....',;;:cllodooolllcc:::;;,''..
       lxxxxxxxkkkkkkkkkkkO0doolllcccccc,
       dxxxxxxxxxxxxxxxxkKWMNkccccccccccc
       xxxxxxxxxxxxxxxOXMMMMMMNkccccccccc,
      .xxxxxxxxxxxxxkXWMMMMMMMMMNkcccccccc
      ,xxxxxxxxxxxOXWMMMMMMMMMMWMWNOlccccc'
      :xxxxxxxxxOXMMMMMMMMMMWWWWWWWWWOlccc:
      lxxxxxxxOXMMMMMMMMMWMWMWWWWWWWWWWOocc.
      cxxxxxOXMMMMMMMMMMMWMWWWWWWWWWWWWWW0o:
      ,xxxOXMMMMMMMMMWWWMWWWWWWWWWWWWWWWWWW0.
      .xxxNMMMMMMMMWWMWWWWWWWWWWWWWWWWWWWWKl
       xx::xXMMMMWWWWWWWWWWWWWWWWWWWWWWWKl.
       xd::::dKWWMWWWWWWWWWWWWWWWWWWWWKl..
       lo::::::l0WWWWWWWWWWWWWWWWWWWKc....
       ;l::::::::lOWWWWWWWWWWWWWWWKl.....
       ':::::::::::ckNWWWWWWWWWWKc......
        ',;:::::::::::dXWWWWWWKc.....
              .....',,;:oOWW0c..
                          .`,

	`        ....',;;:ccloddoolllcc:::;;,'...
       lxxxxxxxxxkkkkkkkkk0xooolllcccccc;
       oxxxxxxxxxxxxxxxk0NMW0occccccccccc
       dxxxxxxxxxxxxxkKWMMMMMMKoccccccccc'
       xxxxxxxxxxxxkKWMMMMMMMMMM0occccccc:
       xxxxxxxxxxkKWMMMMMMMMMMMMMMKdcccccc.
      .xxxxxxxxkKWMMMMMMMMMMMMMMMMMMXdcccc;
      .xxxxxxkKWMMMMMMMMMMMMMMMMMMMMMMXxccc
      .xxxxkXMMMMMMMMMMMMMMMMMMMMMMMMMMMNxc'
       xxkKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNx
       ddKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWO;
       ol:l0MMMMMMMMMMMMMMMMMMMMMMMMMMMWO;.
       cc:::lOWMMMMMMMMMMMMMMMMMMMMMMWk;..
       ;c::::::kNMMMMMMMMMMMMMMMMMMWk,....
       ,:::::::::xNMMMMMMMMMMMMMMWk,.....
       .:::::::::::dKMMMMMMMMMMWx,......
        ',,;:::::::::l0WMMMMMWx'.....
               ....',;;ckNMNd'..
                         .`,

	`        ...'',;;:ccloodooollcc::;;;,'...
       cxxxxxxxxxxxkkkkkkOOdoooolllccccc:
       cxxxxxxxxxxxxxxxOXMMXxcccccccccccc
       cxxxxxxxxxxxxx0NMMMMMMNxcccccccccc.
       lxxxxxxxxxxxONMMMMMMMMMMNxcccccccc:
       lxxxxxxxxx0NMMMMMMMMMMMMMMNklcccccc
       lxxxxxxx0WMMMMMMMMMMMMMMMMMMWklcccc.
       lxxxxx0WMMMMMMMMMMMMMMMMMMMMMMWOlcc:
       cxxxKWMMMMMMMMMMMMMMMMMMMMMMMMMMW0oc
       :x0WMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0'
       ;kWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNd.
       ,;:kNMMMMMMMMMMMMMMMMMMMMMMMMMMMNd'.
       ';;;:dXMMMMMMMMMMMMMMMMMMMMMMMXo'..
       .;;;;;:oKMMMMMMMMMMMMMMMMMMMXl.....
       .;;;;;;;;l0MMMMMMMMMMMMMMMXl......
       .;;;;;;;;;;lOWMMMMMMMMMMKl.......
        .',;;;;;;;;;:kNMMMMMMKc.......
               ....',;:xKMM0:..
                         '`,

	`        ...'',;;:cclloddoollcc::;;;,'...
       cxxxxxxxxxxxxkkkkk0ddoooollllcccc:
       cxxxxxxxxxxxxxxkXWMNOlcccccccccccc
       cxxxxxxxxxxxxkXMMMMMMWOlcccccccccc.
       cxxxxxxxxxxkXMMMMMMMMMMWOlcccccccc,
       :xxxxxxxxOXMMMMMMMMMMMMMMW0occccccc
       :xxxxxxOXMMMMMMMMMMMMMMMMMMMKoccccc
       :xxxxONMMMMMMMMMMMMMMMMMMMMMMMKdccc.
       :xxONMMMMMMMMMMMMMMMMMMMMMMMMMMMXxc;
       :ONMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXd
       :XMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMKl
       .;oKMMMMMMMMMMMMMMMMMMMMMMMMMMMMKc.
       .;;;lOWMMMMMMMMMMMMMMMMMMMMMMMO:...
       .;;;;;ckWMMMMMMMMMMMMMMMMMMWO;.....
       .;;;;;;;:kNMMMMMMMMMMMMMMWk;......
       .;;;;;;;;;:xXMMMMMMMMMMWx,........
        .',,;;;;;;;:oKMMMMMMWx'.......
               ...'',;l0WMXd'..
                        .`,

	`        ...'',;;::clloodoollcc::;;,,'....
       cxxxxxxxxxxxxxkkk0kdddoooolllllllc
       lxxxxxxxxxxxxxx0WMMKdlllllllllllll
       oxxxxxxxxxxxxKWMMMMMMKdlllllllllll.
       dxxxxxxxxxxKWMMMMMMMMMMKdlllllllll'
       dxxxxxxxxKWMMMMMMMMMMMMMMXxlllllll;
       xxxxxxkKWMMMMMMMMMMMMMMMMMMXxlllllc
       xxxxkKWMMMMMMMMMMMMMMMMMMMMMMNkllll
       xxkXMMMMMMMMMMMMMMMMMMMMMMMMMMMNOll.
      .kXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNO'
      .OWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWk.
       ;:kNMMMMMMMMMMMMMMMMMMMMMMMMMMMWx,.
       ,;;:xXMMMMMMMMMMMMMMMMMMMMMMMNd,...
       ';;;;;dXMMMMMMMMMMMMMMMMMMMXo'....
       .;;;;;;;oKMMMMMMMMMMMMMMMXo'......
       .;;;;;;;;;l0WMMMMMMMMMMKl.........
        ..',;;;;;;;cOWMMMMMM0:.........
               ...'',:kNMWO:....
                       .;`,

	`        ...',,;;::cclooddoollcc:;;,,'....
       cxxxxxxxxxxxxxkk0Oddddoooolllllllc
       oxxxxxxxxxxxxxOXMMXkllllllllllllll
       dxxxxxxxxxxxONMMMMMMNkllllllllllll
       xxxxxxxxxxONMMMMMMMMMMNkllllllllll.
      .xxxxxxxx0NMMMMMMMMMMMMMMNOolllllll.
      .xxxxxx0WMMMMMMMMMMMMMMMMMMWOolllll'
      ,xxxx0WMMMMMMMMMMMMMMMMMMMMMMW0olll;
      :xxKWMMMMMMMMMMMMMMMMMMMMMMMMMMW0ol:
      lKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0o
      cKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXo
      .;oKMMMMMMMMMMMMMMMMMMMMMMMMMMMMKl'.
       ;;;l0WMMMMMMMMMMMMMMMMMMMMMMM0c...
       ;;;;;cOWMMMMMMMMMMMMMMMMMMWO:.....
       ';;;;;;ckWMMMMMMMMMMMMMMWk;.......
       .;;;;;;;;:xNMMMMMMMMMMNx,.........
        ..',;;;;;;:dXMMMMMMNd'.........
               ...',;oKMMKo'....
                       '`,

	`         ....',,;;:cclodoollc::;,,'....
       cxxxxxxxxxxxxkkk0Odddddoooooollllc
       lxxxxxxxxxxxxxOXMMXkllllllllllllll
       oxxxxxxxxxxxONMMMMMMNkllllllllllll
       dxxxxxxxxxONMMMMMMMMMMNkllllllllll
       xxxxxxxxONMMMMMMMMMMMMMMNkllllllll.
       xxxxxxONMMMMMMMMMMMMMMMMMMNkllllll.
      .xxxx0NMMMMMMMMMMMMMMMMMMMMMMNOllll.
      .xx0NMMMMMMMMMMMMMMMMMMMMMMMMMMWOll'
      'ONMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWO;
      .0WMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0,
       ;cOWMMMMMMMMMMMMMMMMMMMMMMMMMMMO:.
       ;;;ckNMMMMMMMMMMMMMMMMMMMMMMWk;...
       ';;;;:xNMMMMMMMMMMMMMMMMMMWx,.....
       .;;;;;;:xNMMMMMMMMMMMMMMNx,.......
       .;;;;;;;;:dXMMMMMMMMMMXd'.........
         ..',;;;;;;oKMMMMMMXo'........
                ...',l0MMKl....
                       .`,

	`           ....',;;:clooollc:;;,'....
       ;dxxxxxxxxxkkkkk0Oddddddooooooool:
       cxxxxxxxxxxxxxOXMMXklllllllllllllc
       lxxxxxxxxxxxONMMMMMMNkllllllllllll
       lxxxxxxxxxkXMMMMMMMMMMXxllllllllll
       oxxxxxxxOXMMMMMMMMMMMMMMXxllllllll
       dxxxxxkXMMMMMMMMMMMMMMMMMMXxllllll
       xxxxOXMMMMMMMMMMMMMMMMMMMMMMXxllll
       xxkXMMMMMMMMMMMMMMMMMMMMMMMMMMNxll.
       kXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXx.
       xNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWx.
       ;:xNMMMMMMMMMMMMMMMMMMMMMMMMMMWx,.
       ';;:dXMMMMMMMMMMMMMMMMMMMMMMNd,...
       .;;;;;oXMMMMMMMMMMMMMMMMMMXo'.....
       .;;;;;;;oXMMMMMMMMMMMMMMXo'.......
        ;;;;;;;;;oKMMMMMMMMMMKl'.........
          ..',;;;;;l0MMMMMMKc........
                ...',c0WM0c...
                       .`,

	`             ...',;;:cloolc:;,,'...
       ,lodxxxkkkkkkkkk0Oddddddddoooolc:,
       :xxxxxxxxxxxxxOXMMXklllllllllllll:
       :xxxxxxxxxxxkXMMMMMMNxlllllllllllc
       :xxxxxxxxxkXMMMMMMMMMMXdlllllllllc
       cxxxxxxxkXMMMMMMMMMMMMMMKxlllllllc
       cxxxxxkXMMMMMMMMMMMMMMMMMMXdlllllc
       cxxxkKWMMMMMMMMMMMMMMMMMMMMMKdlllc
       cxxKMMMMMMMMMMMMMMMMMMMMMMMMMMKdlc
       lKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0o
       cKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXl
       .;oKMMMMMMMMMMMMMMMMMMMMMMMMMMXl'.
       .;;;oKMMMMMMMMMMMMMMMMMMMMMMKl'...
       .;;;;;l0MMMMMMMMMMMMMMMMMMKc......
       .;;;;;;;l0MMMMMMMMMMMMMMKc........
        ;;;;;;;;;l0WMMMMMMMMM0c.........
           ..',;;;;cOWMMMMM0:.......
                 ...':OWWO:...
                       .`,

	`              ...',;:clollc:;,'...
       .:cldxkkkkkkkkkk0Oddddddddddolc:;'
       ;xxxxxxxxxxxxxOXMMXxlllllllllllll:
       ;xxxxxxxxxxxkXMMMMMMXxlllllllllll:
       ;xxxxxxxxxkKMMMMMMMMMMKdlllllllll;
       ,xxxxxxxxKWMMMMMMMMMMMMMKdlllllll;
       ,xxxxxxKWMMMMMMMMMMMMMMMMM0olllll;
       ,xxxx0WMMMMMMMMMMMMMMMMMMMMW0olll,
       'xx0WMMMMMMMMMMMMMMMMMMMMMMMMW0ol,
       'ONMMMMMMMMMMMMMMMMMMMMMMMMMMMMWO,
       .OWMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0'
       .;cOWMMMMMMMMMMMMMMMMMMMMMMMMMO:..
       .;;;cOWMMMMMMMMMMMMMMMMMMMMWO:....
        ;;;;;cOWMMMMMMMMMMMMMMMMMO;......
        ;;;;;;;cOWMMMMMMMMMMMMMO:.......
        ,;;;;;;;;cOWMMMMMMMMWO:.........
           ...';;;;ckWMMMMWO;.......
                  ..';kWWk,..
                       .`,

	`                ..',;:lolc:;,'..
       .,;cloxkkkkkkkkk0Odddddddddolc;,'.
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll:
       ;xxxxxxxxxxxkXMMMMMMXdlllllllllll;
       ,xxxxxxxxxxKWMMMMMMMMMKolllllllll;
       ,xxxxxxxxKWMMMMMMMMMMMMW0olllllll,
       'xxxxxx0WMMMMMMMMMMMMMMMMWOllllll,
       'xxxx0NMMMMMMMMMMMMMMMMMMMMWOllll,
       'xxONMMMMMMMMMMMMMMMMMMMMMMMMNkll'
       .kXMMMMMMMMMMMMMMMMMMMMMMMMMMMMXx'
       .xNMMMMMMMMMMMMMMMMMMMMMMMMMMMMNx.
       .;:xNMMMMMMMMMMMMMMMMMMMMWMWMWx,..
       .;;;:xNMMMMMMMMMMMMMMMWWWWMNx,....
       .;;;;;:xNMMMMMWMMMWMWWWWWWx,......
        ;;;;;;;:xNMWWWWMWWWWWWWx,........
        ',;;;;;;;:xNWWWWWWWWNx,.........
            ...,;;;:xNWWWWWx,......
                  ...,xNNx'.
`,

	`                 ..';:colc:,'...
       ..,;:loxkkkkkkkk0Oddddddddolc;,'..
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll:
       ;xxxxxxxxxxxkXMMMMMMXdlllllllllll;
       ,xxxxxxxxxxKWMMMMMMMMM0olllllllll;
       ,xxxxxxxx0WMMMMMMMMMMMMWOolllllll,
       ,xxxxxxONMMMMMMMMMMMMMMMMWkllllll'
       'xxxxONMMMMMMMMMMMMMMMMMMMMNxllll'
       'xxkXMMMMMMMMMMMMMMMMMMMMMMMMXdll.
       .xKWMMMMMMMMMMMMMMMMMMMMMMMMMMMKo.
       .dKWWWWWWWWWWWWWWWWWWWWWWWWWWWWKl.
       .:;oKWWWWWWWWWWWWWWWWWWWWWWWWXl'..
       .:;;;oKWWWWWWWWWWWWWWWWWWWWXo'....
       .;;;;;;oXWWWWWWWWWWWWWWWWXo'......
       .;;;;;;;;dXWWWWWWWWWWWWNd'........
        .',;;;;;;:dXWWWWWWWWNd'.........
             ..',;;:dXWWWWNd'.....
                   ..'oXXd.
`,

	`                 ..',:col:;,'..
        ..';:loxkkkkkkk0Odddddddolc;,'..
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll:
       ;xxxxxxxxxxxxKMMMMMMKdlllllllllll;
       ;xxxxxxxxxx0WMMMMMMMMW0llllllllll,
       ,xxxxxxxxONMMMMMMMMMMMMWkllllllll,
       ,xxxxxxkNMMMMMMMMMMMMMMMMNxllllll'
       'xxxxkXMMMMMMMMMMMMMMMMMMMMXdllll'
       'xxxKWMMMMMMMMMMMMMMMMMMMMMMM0olc.
       .xOWMMMMMMMMMMMMMMMMMMMMMMMMMMWOc.
       .dOWWWWWWWWWWWWWWWWWWWWWWWWWWWWO:.
       .l;cOWWWWWWWWWWWWWWWWWWWWWWWWO:.'.
       .c;;;lOWWWWWWWWWWWWWWWWWWWW0c.....
       .:;;;;;l0WWWWWWWWWWWWWWWW0c.......
       .;;;;;;;;oKWWWWWWWWWWWWKl.........
        ..',;;;;;;oKWWWWWWWWKl'........
              ..,;;;oKWWWWXo'.....
                   ..'lKKl.
`,

	`                  ..,;clc:;'..
         ..';:ldxkkkkkk0Odddddddlc:,'..
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll;
       ;xxxxxxxxxxxxKMMMMMMKolllllllllll;
       ;xxxxxxxxxx0WMMMMMMMMWOllllllllll,
       ,xxxxxxxxONMMMMMMMMMMMMNkllllllll'
       ,xxxxxxkXMMMMMMMMMMMMMMMMXdllllll'
       'xxxxxKWMMMMMMMMMMMMMMMMMMM0olllc.
       'xxxOWMMMMMMMMMMMMMMMMMMMMMMWOllc.
       'xkXMMMMMMMMMMMMMMMMMMMMMMMMMMNx:.
       .xdXWWWWWWWWWWWWWWWWWWWWWWWWWWNd;.
       .o;:xNWWWWWWWWWWWWWWWWWWWWWWNx,.,.
       .l;;;:kNWWWWWWWWWWWWWWWWWWNk;...'.
       .:;;;;;:kNWWWWWWWWWWWWWWNk;.......
       .;;;;;;;;cOWWWWWWWWWWWW0:.........
         ..';;;;;;l0NWWWWWWW0c.........
              ..',;;l0WWWWKl.....
                    ..c00c.
`,

	`                  ..';:lc:,'..
          ..';codkOkkkk0Oddddddoc:,'..
       'odxxxxxxxxxxxkXMMXxllllllllllllc,
       ;xxxxxxxxxxxxKMMMMMMKolllllllllll;
       ;xxxxxxxxxxOWMMMMMMMMWkllllllllll,
       ,xxxxxxxxkXMMMMMMMMMMMMXxllllllll'
       ,xxxxxxxKMMMMMMMMMMMMMMMMKolllllc'
       ,xxxxx0WMMMMMMMMMMMMMMMMMMWOllllc.
       'xxxkNMMMMMMMMMMMMMMMMMMMMMMNxll:.
       'xxKMMMMMMMMMMMMMMMMMMMMMMMMMMKo;.
       .xo0WWWWWWWWWWWWWWWWWWWWWWWWWWKc;.
       .x;;oKNNNNNNNNNNNNNNNNNNNNNNXl'.,.
       .o;;;;dKNNNNNNNNNNNNNNNNNNXo'...'.
       .c;;;;;:xXNNNNNNNNNNNNNNNd'.....'.
       .;;;;;;;;:kNNNNNNNNNNNNk;.........
         ...,;;;;;cONNNNNNNNO:........
               ..,;;cONNNN0:....
                    ..:OO:
`,

	`                   ..,:lc;,..
           ..,:coxOOkkk0Oddddddl:;'..
       .coxxxxxxxxxxxkXMMXdlllllllllllc;'
       ;xxxxxxxxxxxxKWMMMMM0olllllllllll;
       ;xxxxxxxxxxONMMMMMMMMWkllllllllll,
       ;xxxxxxxxxXMMMMMMMMMMMMXdlllllllc'
       ,xxxxxxx0WMMMMMMMMMMMMMMW0llllllc'
       ,kxxxxkNMMMMMMMMMMMMMMMMMMNxllll:.
       'kxxxKMMMMMMMMMMMMMMMMMMMMMMKoll;.
       'kxOWMMMMMMMMMMMMMMMMMMMMMMMMWOc;.
       .kdkNWWWWNWNWNWNWWNWNWNWNWNWNNk;;.
       .k:;cONNNNNNNNNNNNNNNNNNNNNN0:..;.
       .d;;;;l0NNNNNNNNNNNNNNNNNN0c....,.
       .c;;;;;;oKNNNNNNNNNNNNNNXl'.....'.
       .;;;;;;;;:dXNNNNNNNNNNNd'........
          ..',;;;;:xXNNNNNNNx;.......
               ..',;cONNNNO;....
                    ..;kk;
`,

	`                   ..,:lc;'..
            ..,:ldkOkkkKOdddddoc;,..
       .;coxxxxxxxxxxkXMMXdllllllllolc:,.
       ;xxxxxxxxxxxx0WMMMMM0llllllllllll;
       ;xxxxxxxxxxkNMMMMMMMMNxllllllllll,
       ;xxxxxxxxxKMMMMMMMMMMMMKolllllllc'
       ,kxxxxxxOWMMMMMMMMMMMMMMWkllllll:'
       ,kxxxxkXMMMMMMMMMMMMMMMMMMXdllll;.
       'kxxx0WMMMMMMMMMMMMMMMMMMMMM0llc;.
       'kkkNMMMMMMMMMMMMMMMMMMMMMMMMNx:;.
       'kxdXNNNNNNNNNNNNNNNNNNNNNNNNXd;;.
       .kl;:xXNNNNNNNNNNNNNNNNNNNNNx,.';.
       .x;;;;ckNNNNNNNNNNNNNNNNNNk;....,.
       .l;;;;;;c0NNNNNNNNNNNNNN0:......'.
        ,;;;;;;;;oKNNNNNNNNNNXo'........
           ..,;;;;:dXNNNNNNXd,.......
                ..,;:kXNNNk,....
                     .,xx,
`,

	`                    .';c:;'.
             .';coxOOkkKOdddddl:,..
       .,;coxxxxxxxxxkXMMXdlllllllooc:,'.
       ;xxxxxxxxxxxx0WMMMMM0llllllllllll;
       ;xxxxxxxxxxkXMMMMMMMMNdllllllllll'
       ;xxxxxxxxx0WMMMMMMMMMMM0llllllllc'
       ,kxxxxxxkNMMMMMMMMMMMMMMNxllllll:'
       ,kxxxxxKWMMMMMMMMMMMMMMMMMKolllc;.
       ,kkxxkNMMMMMMMMMMMMMMMMMMMMWkll:;.
       'kkxKMMMMMMMMMMMMMMMMMMMMMMMMKo;;.
       'kko0NNNNNNNNNNNNNNNNNNNNNNNN0c;;.
       .ko;;oKNNNNNNNNNNNNNNNNNNNNKo'.';.
       .k:;;;:xXNNNNNNNNNNNNNNNNXd,....;.
       .l;;;;;;:kXNNNNNNNNNNNNNk;......'.
        ',;;;;;;;l0NNNNNNNNNN0c.........
           ..',;;;;oKNNNNNNKo'......
                ..';:xXNNXx,...
                     .'dd
`,

	`                    .';c:,..
              .,:lxOOkkK0ddddoc;'.
        .,:ldxxxxxxxxxXMMXdllllllool:;'.
       ;xxxxxxxxxxxxOWMMMMWOllllllllllll;
       :xxxxxxxxxxxXMMMMMMMMXollllllllll'
       ;kxxxxxxxxOWMMMMMMMMMMWOllllllllc'
       ;kxxxxxxxXMMMMMMMMMMMMMMXdllllll:'
       ,kkxxxx0WMMMMMMMMMMMMMMMMWOllllc;.
       ,kkxxxXMMMMMMMMMMMMMMMMMMMMXdll:;.
       'OkkOWMMMMMMMMMMMMMMMMMMMMMMWOc;;.
       'kOokXNNNNNNNNNNNNNNNNNNNNNNNk;;;.
       .kx;;cOXXXXXXXXXXXXXXXXXXXX0:..,;.
       .k:;;;;o0XXXXXXXXXXXXXXXXKl'....;.
       .l;;;;;;:dXXXXXXXXXXXXXXd'......'.
        .';;;;;;;ckXXXXXXXXXXO;.........
            ..';;;;l0XXXXXX0l'.....
                 ..,;dKXXXd'...
                     ..oo
`,

	`                    ..;c:,..
              ..;ldkOkkK0ddddl:,..
        ..,:lxkxxxxxxxXMMXdllllllolc;'..
       ;xxxxxxxxxxxxOWMMMMWOllllllllllll;
       :xxxxxxxxxxxKMMMMMMMMKollllllllll'
       ;kxxxxxxxxONMMMMMMMMMMWkllllllllc'
       ;OxxxxxxxKMMMMMMMMMMMMMMKollllll;'
       ,OkxxxxkNMMMMMMMMMMMMMMMMWxllll:;.
       ,OOxxxKMMMMMMMMMMMMMMMMMMMMKllc;;.
       'OOkkNMMMMMMMMMMMMMMMMMMMMMMNd:;;.
       'OOxdKNNNNNNNNNNNNNNNNNNNNNNXo,;;.
       .Ok:;:xXXXXXXXXXXXXXXXXXXXXx,..;;.
       .kc;;;;cOXXXXXXXXXXXXXXXXO:....';.
       .l;;;;;;;oKXXXXXXXXXXXXKl'......'.
        ..,;;;;;;:xXXXXXXXXXXx,........
             .',;;;cOXXXXXXO:......
                 ..,;oKXXKl'..
                      .ll
`,

	`                     .,c;,.
               .,:oxOOkK0oddoc;'.
         .';coxkxxxxxxXMMXolllllooc:,..
       ,dxxxxxxxxxxxOWMMMMWkllllllllllll,
       :xxxxxxxxxxxKMMMMMMMMKlllllllllll'
       ;kxxxxxxxxkNMMMMMMMMMMNxllllllll:'
       ;Oxxxxxxx0WMMMMMMMMMMMMM0llllllc;'
       ,OOxxxxxXMMMMMMMMMMMMMMMMXdllll:;.
       ,OOkxxOWMMMMMMMMMMMMMMMMMMWkllc;;.
       'OOOxKMMMMMMMMMMMMMMMMMMMMMMKo:;;.
       'OOkoONNNNNNNNNNNNNNNNNNNNXN0:;;;.
       .OOc;;oKXXXXXXXXXXXXXXXXXXKo'.';;.
       .Oc;;;;:xXXXXXXXXXXXXXXXXx,....';.
       .c;;;;;;;cOXXXXXXXXXXXX0:.......'.
         .',;;;;;;dKXXXXXXXXXd'........
             ..,;;;ckXXXXXXk;.....
                  .';l0XXKc...
                      .cl
`,

	`                     .,c;'.
               ..;lxOOkK0oddoc,..
          .':ldkxxxxxxXMMXollllool:,'.
       'oxxxxxxxxxxxkWMMMMWxlllllllllllc'
       :xxxxxxxxxxx0WMMMMMMM0lllllllllll'
       ;OxxxxxxxxxXMMMMMMMMMMXdllllllll:'
       ;OkxxxxxxOWMMMMMMMMMMMMWkllllllc;'
       ,OOxxxxxKMMMMMMMMMMMMMMMMKollll:;.
       ,OOOxxkNMMMMMMMMMMMMMMMMMMNdll:;;.
       ,OOOkOWMMMMMMMMMMMMMMMMMMMMWOc;;;.
       'OOOoxXXXXXXXXXXXXXXXXXXXXXXk,;;;.
       .OOl;;cOXXXXXXXXXXXXXXXXXX0:..';;.
       .Ol;;;;;dKXXXXXXXXXXXXXXKo'....';.
       .c;;;;;;;:kXXXXXXXXXXXXk;.......'.
         ..';;;;;;o0XXXXXXXXKl'.......
              .',;;:xKXXXXXx,.....
                  ..,cOXX0:...
                      .:c
`,

	`                     .,:;'.
                .,cdkOkK0oddl:'.
           .,coxkxxxxxKMMXollllooc;'.
       .coxxxxxxxxxxkNMMMMWxllllllllllc;.
       :xxxxxxxxxxxOWMMMMMMMOlllllllllll,
       :OxxxxxxxxxKMMMMMMMMMMKollllllll:'
       ;OkxxxxxxkNMMMMMMMMMMMMWxllllllc;'
       ;OOkxxxxOWMMMMMMMMMMMMMMMOllllc;;.
       ,OOOxxxKMMMMMMMMMMMMMMMMMMXoll:;;.
       ,OO0OkNMMMMMMMMMMMMMMMMMMMMNd:;;;.
       'O0OxoKXXXXXXXXXXXXXXXXXXXXKl,;;;.
       'OOd;;:xKXXXXXXXXXXXXXXXXXk,..';;.
       .Ol;;;;;l0XXXXXXXXXXXXXX0c.....';.
       .:;;;;;;;:dKXXXXXXXXXXKd'.........
          ..,;;;;;cOXXXXXXXX0:........
              ..,;;:dKXXXXKd'....
                  ..,cOXXO;..
                      .;:
`,

	`                     .':;..
                .':okOkK0odoc;..
           .':lxkxxxxxKMMXollllol:,..
       .;ldxxxxxxxxxkNMMMMNdllllllllll:,.
       :xxxxxxxxxxxOWMMMMMMWklllllllllll,
       :Oxxxxxxxxx0MMMMMMMMMM0lllllllll:'
       ;0OxxxxxxxXMMMMMMMMMMMMNollllllc;'
       ;00kxxxxkNMMMMMMMMMMMMMMWxllllc;;.
       ,000kxx0WMMMMMMMMMMMMMMMMMOllc;;;.
       ,000OxKMMMMMMMMMMMMMMMMMMMMKl:;;;.
       '000klOXXXXXXXXXXXXXXXXXXXX0:;;;;.
       '00d:;;oKKKKKKKKKKKKKKKKKKo'..,;;.
       .Ol;;;;;ckKKKKKKKKKKKKKKk;.....';.
        ,;;;;;;;;o0KKKKKKKKKKKl'........
          ..';;;;;:kKKKKKKKKO;.......
              ..';;;o0KKKKKo'....
                   .':kKKk,..
                       ;
`,

	`                     .':,..
                 .;lxOkK0oooc,.
            .,cdkkxxxxKMMKolllloc;'.
        ':oxxxxxxxxxxNMMMMNdlllllllllc,.
       :xxxxxxxxxxxkWMMMMMMWxlllllllllll,
       :OxxxxxxxxxOWMMMMMMMMMOlllllllll:'
       ;0OxxxxxxxKMMMMMMMMMMMMKlllllllc:'
       ;00OxxxxxXMMMMMMMMMMMMMMNdllllc::'
       ,000kxxkWMMMMMMMMMMMMMMMMWxllc:::.
       ,0000kOWMMMMMMMMMMMMMMMMMMMOc::::.
       '000OoxKXXXXXXXXXXXXXXXXXXXx,;:::.
       '00x:;;cOKKKKKKKKKKKKKKKK0:...,;:.
       .Ol;;;;;:dKKKKKKKKKKKKKKd,.....';.
        .,;;;;;;;cOKKKKKKKKKK0c.........
           .',;;;;:xKKKKKKKKx,.......
               ..,;;lOKKKK0c.....
                   .';xKKx,..
                       ,
`,

	`                      .;;'.
                 .;ldkkXxoddl;.
            .,:oxxxxxkNMWkllllool:,.
        ':oxxxxxxxxxOWMMMM0llllllllllc;.
       ;xxxxxxxxxxxKMMMMMMMKolllllllllll;
       ;OxxxxxxxxxXMMMMMMMMMNdlllllllllc,
       ,OkxxxxxxkWMMMMMMMMMMMWxlllllllc:,
       ,OOkxxxx0WMMMMMMMMMMMMMMOlllllc::'
       'OOOxxxKMMMMMMMMMMMMMMMMMKlllc:::'
       'O0OOkNMMMMMMMMMMMMMMMMMMMXoc::::'
       .O0Oxl0XXXXXXXXXXXXXXXXXXX0:;::::'
       .OOo;;;dKKKKKKKKKKKKKKKKKd'''';::.
       .kc;;;;;cOKKKKKKKKKKKKK0c'''''',:.
        .';;;;;;;dKKKKKKKKKKKx,'''''''..
           ..,;;;;cOKKKKKKK0l'...'...
               .',;:dKKKKKk;''...
                   .'cOKKl...
                      .;
`,

	`                      .,;,.
                 .,cokOKddddl:'.
            .,:lxxxdd0WMXdllloodo:,.
        ':lxdddddddxXMMMMWxllllllllloc;.
       'xdddddddddkNMMMMMMWkllllllllllll;
       'kddddddddOWMMMMMMMMMOllllllllllc;
       'OxdddddxKMMMMMMMMMMMMKllllllllcc,
       .OOxdddkNMMMMMMMMMMMMMMXolllllccc,
       .OOkddOWMMMMMMMMMMMMMMMMNdlllcccc,
       .OOOk0MMMMMMMMMMMMMMMMMMMWxlccccc'
       .OOOlxXXXXXXXXXXXXXXXXXXXXd;:cccc'
       .Ok:;;cOKKKKKKKKKKKKKKKK0:''';:cc'
       .x:;;;;;dKKKKKKKKKKKKKKx,'''''';:.
        .';;;;;;cOKKKKKKKKKK0c'''''''''..
           ..,;;;;dKKKKKKKKk;''''''..
               .',;cOKKKKKo'''...
                   .,dKKO;..
                      .
`,

	`                      .';;.
                 .':lxKkdddxo:'.
            .':ldxddxXMWOollooodoc,.
        ':lddddddddOWMMMMKollllllllooc;.
       .xddddddddd0MMMMMMMXollllllllllll;
       .kdddddddxXMMMMMMMMMNdllllllllllc;
       .OxdddddkWMMMMMMMMMMMWdllllllllcc;
       .Okdddd0MMMMMMMMMMMMMMWkllllllccc,
       .OOxdxXMMMMMMMMMMMMMMMMMOllllcccc,
       .OOkkWMMMMMMMMMMMMMMMMMMM0llccccc,
        OOdl0XXXXXXXXXXXXXXXXXXX0::ccccc'
        Od;;;dKXXXXXXXXXXXXXXXKd''',:ccc'
        d;;;;;cOXXXXXXXXXXXXX0c'''''',:c'
        ..,;;;;;oKXXXXXXXXXXk,'''''''''..
            .,;;;:OXXXXXXXKo'''''''..
               .',;oKXXXXO:'''...
                  ..;kXKo'..
                      ,
`,

	`                       .,;'.
                 ..;cxXdddxxdc,.
            .';coxddOWMNdooooodddc,.
        ':lddddddddKMMMMWxoooooooooooc,.
       .xddddddddxNMMMMMMWkooooooooooooo:
       .xdddddddOWMMMMMMMMMOooooooooooll;
        kddddddKMMMMMMMMMMMM0oooooooolll;
        kxdddkNMMMMMMMMMMMMMMKoooooollll,
        kkddOWMMMMMMMMMMMMMMMMXooollllll,
        kkxKMMMMMMMMMMMMMMMMMMMXolllllll,
        kkcxKXXXXXXXXXXXXXXXXXXXl:llllll,
        kl;;cOXXXXXXXXXXXXXXXXO:'',;clll'
        o;;;;;oKXXXXXXXXXXXXKx,'''''';cl'
        ..,;;;;:kXXXXXXXXXXKc'''''''''''.
            .,;;;oKXXXXXXXO;'''''''..
               .';:kKXXXKd,'''...
                  ..c0XO:'..
                     .;
`,

	`                       .';,.
                  .,:x0ddxxkxl,.
            ..;coddxKMM0oooooddddl;.
        ':ldddddddkNMMMMKooooooooooodc,.
        dddddddddOWMMMMMMXoooooooooooooo;
        xddddddxXMMMMMMMMMNdoooooooooooo;
        xdddddkWMMMMMMMMMMMNdooooooooolo;
        xdddd0MMMMMMMMMMMMMMWxooooooooll,
        xxdxNMMMMMMMMMMMMMMMMWxoooololll,
        dxkWMMMMMMMMMMMMMMMMMMWxoololllo,
        ddcOXXXXXXXXXXXXXXXXNNNk:lololol'
        o:;;oKXXXXXXXXXXXXXXXKo''';clolo'
        :;;;;:kXXXXXXXXXXXXX0:'''''',:lo'
         .';;;;l0XXXXXXXXXXk,'''''''''',.
            .';;;xXXXXXXXKo''''''''..
               .',l0XXXX0c''''...
                  .'dKXd,'..
                     .
`,

	`                       ..,;.
                  .':kxdxxxkko;.
            ..,:lddkNMWxoooodddxxl;.
        ';lddddddd0MMMMWxooooooooooddc,.
        odddddddxXMMMMMMMkoooooooooooooo;
        oddddddOWMMMMMMMMMOooooooooooooo;
        odddddKMMMMMMMMMMMMOoooooooooooo,
        odddkNMMMMMMMMMMMMMMOooooooooooo,
        ldd0MMMMMMMMMMMMMMMMM0oooooooooo,
        lxXMMMMMMMMMMMMMMMMMMM0ooooooooo'
        cloKXXXXXXXXXXXXXNNNNNKclooooooo'
        c;,:kXXXXXXXXXXXXXXXXk;,,,:loooo'
        ,,,,,l0XXXXXXXXXXXXXd,,,,,,,;coo.
         .',,,;xXXXXXXXXXXKc,,,,,,,,,,,:.
            .',,c0XXXXXXX0:,,,,,,,'...
               .';dKXXXKd,,,,'...
                  .;OXKc'...
                     ,
`,

	`                       ..,;'
                  ..:koxxxkkkd:.
            ..,:lod0WMKoooodddxxxl;.
        .;codddddxXMMMMKoooooooooodddc,.
        :dddddddkWMMMMMMXdoooooooooooooo,
        :ddddddKMMMMMMMMMXdooooooooooood,
        :ddddkNMMMMMMMMMMMNdoooooooooddd,
        :ddd0MMMMMMMMMMMMMMNdooooooodddd'
        ;dxNMMMMMMMMMMMMMMMMNdoooodddddd'
        ;kWMMMMMMMMMMMMMMMMMMNdodddddddd'
        ;:kXXXXXXXXXXXXNNNNNNNocoddddddd.
        ',,l0XXXXXXXXXXXXXXXKc,,,:lddddd.
        .,,,;xKXXXXXXXXXXXXO:,,,,,,;:ldd.
         ..,,,cOXXXXXXXXXXk,,,,,,,,,,,;c.
            .',;dKXXXXXXXd,,,,,,,,,...
               .':kXXXX0c,,,,'...
                  .cKXk;'...
                    .;
`,

	`                       ..';,.
                  .'oddxxkkkkxc'
            ..,:ldxXMWkooodddxxxxo;.
        .;codddddOWMMMWkoooooooooddddc'.
        ,dddddddKMMMMMMWkooooooooooooood'
        ,dddddkNMMMMMMMMWOoooooooooooodd'
        'dddd0MMMMMMMMMMMMkooooooooodddd'
        'ddxNMMMMMMMMMMMMMWkoooooodddddd.
        'd0WMMMMMMMMMMMMMMMWkooooddddddd.
        .KWWMMMMMMMMMMMMMMMMWxoddddddddd.
        .l0XXXXXXXXXXXNNNNNNNOcodddddddd.
        .,;dXXXXXXXXXXXXXXXXx,,,;cdddddd.
         ,,,:OXXXXXXXXXXXXKo,,,,,,,:lddd.
          .,,,oKXXXXXXXXXKc,,,,,,,,,,,:l
            ..,;kXXXXXXXO:,,,,,,,,,'..
               .'l0XXXXx,,,,,'...
                  .dXXo,'...
                    .
`,

	`                       ..';;.
                  .;xoxxkkkkkxc'
            ..,:ookWMNddddddxxxkko;.
         ;cooooodKMMMMXdddododdodddxdc'
        .doooooxNMMMMMMKdddoododoododood.
        'doood0MMMMMMMMMXdddoddodooddddx.
        ,oooxNMMMMMMMMMMMXdoodoodooddxxx.
        ,od0WMMMMMMMMMMMMMKdodoooddxxxxx.
        ;xXMMMMMMMMMMMMMMMMKoooddxxxxxxx.
        cWWWMMMMMMMMMMMMMMMM0ddxxxxxxxxx.
        .dKXXXXXXXXXXNNNNNNNKcoxxxxxxxxx
        .,:OXXXXXXXXXXXXXXX0;,,;coxxxxxx
         ,,,oKXXXXXXXXXXXXk;,,,,,,;coxxx
          .',;xXXXXXXXXXXx,,,,,,,,,,,;co
             .,c0XXXXXXKo,,,,,,,,,,'..
               .'dKXXXKc,,,,,'...
                  'OX0:,'..
                    ,
`,

	`                       ..';:.
                 ..cxoxkkkOOOkl,.
            ..;codKMM0dddddxxxkkkd:.
         ;coooooxNMMMMOdddddddddddxxd:.
        .oooooo0MMMMMMWkdddddddddddddddd
        'ooooxNMMMMMMMMWxdddddddddddddxx.
        ,oooOWMMMMMMMMMMWxddddddddddxxxx
        :oxXMMMMMMMMMMMMMWxdddddddxxxxxx
        lOWMMMMMMMMMMMMMMMNdddddxxxxxxxx
        0WWWWMMMMMMMMMMMMMMXddxxxxxxxxxx
        ,kXXXXXXXXXXXNNNNNNNooxxxxxxxxxx
        .,lKXXXXXXXXXXXXXXKl,,;coxxxxxxx
         ',;xXXXXXXXXXXXXKc,,,,,,;:ldxxx
          ..,:OXXXXXXXXX0:,,,,,,,,,,,:ld
             .,dKXXXXXXO;,,,,,,,,,,'..
               .;kXXXXk;,,,,,'...
                  :0Xx,,'..

`,

	`                       ..',:,
                 .'dddkkkOOOOOo;.
            .';coxNMWxdddddxxkkkOd:.
         ,cooooo0WMMMNddddddddddddxxd:.
         ooooodXMMMMMMXddddddddddddddddd
        .ooooOWMMMMMMMM0dddddddddddddxkk
        ;oodXMMMMMMMMMMM0ddddddddddxkkkx
        loOWMMMMMMMMMMMMMOdddddddxkkkkkx
        dXMMMMMMMMMMMMMMMMkddddxkkkkkkkx
       'NWWWWWMMMMMMMMMMMMWxdxkkkkkkkkkd
        c0XXXXXXXXXXNNNNNNNkoxkkkkkkkkkd
        .,dXXXXXXXXXXXXXXXx;,,:lxkkkkkko
         .':OXXXXXXXXXXXXx,,,,,,,:ldkkko
           .,lKXXXXXXXXXo,,,,,,,,,,,;col
             .;kXXXXXXXl,,,,,,,,,,,'...
               .c0XXXKc,,,,,,'...
                 .oXKc,'...
                   .
`,

	`                       ..',;;.
                 .;koxkkOOOOOOd;.
            .';coOWMXdddddxxxkkOOd:.
         'coooodXMMMM0dddddddddddxxkd;.
         looookWMMMMMMkddddddddddddddddl
        .ooodKMMMMMMMMWxdddddddddddddxko
        :ookWMMMMMMMMMMNdddddddddddxkkko
        odKMMMMMMMMMMMMMXdddddddxkkkkkko
       .kWMMMMMMMMMMMMMMMKddddxkkkkkkkkl
       oWWWWWWMMMMMMMMMMMMOdxkkkkkkkkkkl
        oKXXXXXXXXXXXNNNNNKlxkkkkkkkkkkc
         ;kXXXXXXXXXXXXXX0:;,:ldkkkkkkkc
         .'l0XXXXXXXXXXXO:,;,;,;;coxkkk:
           .,dXXXXXXXXXk;;,;,;,;,;,;:lx:
             .:0XXXXXXO;;;,;,;,;,;,'...
               .oKXXXk;;;,;,,'...
                 .xXk;,'...
                   '
`,

	`                       ..',;:.
                 .cxdxkOOOOOOOx:.
            .';coKMMOddddxxxkkOOOd:.
         .:loookWMMMWxddddddddddxxkkd;.
         cooooKMMMMMMXddddddddddddddddx;
        .oooxNMMMMMMMM0dddddddddddddxkOc
        :ooKMMMMMMMMMMMkddddddddddxkOOOc
       .okNMMMMMMMMMMMMWxdddddddxkOOOOOc
       :KMMMMMMMMMMMMMMMNxddddkOOOOOOOOc
       KWWWWWWMMMMMMMMMMMXdxkOOOOOOOOOO:
       .xXXXXXXXXXXXXNNNNNodkOOOOOOOOOO;
         :0XXXXXXXXXXXXXXl;;;cdkOOOOOOO;
          'oKXXXXXXXXXXKl;;;;;;;coxOOOO,
           .;kXXXXXXXXKc;;;;;;;;;;;;ldk'
             .lKXXXXXKl;;;;;;;;;;;;,...
               .xXXXXl;;;;;;,'...
                 ,0Xo;,'...

`,

	`                       ..',;:'
                .'dddkOOOOOOOOkc.
            .';lxNMNxdddxxxkkOOOOx;.
          :lllo0MMMMKddddddddddxxxkko,
         ;llldNMMMMMMOddddddddddddddddx.
        .llo0WMMMMMMMWxdddddddddddddkOO;
        :lxNMMMMMMMMMMXddddddddddxkOOOOc
       .o0WMMMMMMMMMMMMOdddddddxkOOOOOOl
       oNMMMMMMMMMMMMMMWkddddxOOOOOOOOOl
      .NWWWWWWMMMMMMMMMMNxxkOOOOOOOOOOO:
       .kXXXXXXXXXXXXNNNNxdkOOOOOOOOOOO;
         lKXXXXXXXXXXXXXd;;;cokOOOOOOOO'
          'xXXXXXXXXXXXd;;;;;;;:ldkOOOO.
            :0XXXXXXXXx;;;;;;;;;;;;coxO.
             .dXXXXXXO;;;;;;;;;;;;;,...
               'OXXXO;;;;;;;,'...
                 :K0:;,'..
                  .
`,

	`                       ...,;:,
                .,koxkOOOOOOOOOl.
            .';lOWMKxxxxxxkkkOOOOx;.
          ;llldXMMMMkxxxxxxxxxxxxkkkl'
         'lllkWMMMMMNxxxxxxxxxxxxxxxxxd
         lldXMMMMMMMM0xxxxxxxxxxxxxxkO0'
        :lOWMMMMMMMMMWkxxxxxxxxxxxOO000;
       'dXMMMMMMMMMMMMNxxxxxxxxkO000000:
       OWMMMMMMMMMMMMMM0xxxxxkO00000000:.
      :NWWWWWWMMMMMMMMMWkxxOO0000000000;
       ,OXXXXXXXXXXXXNNNkdkO0000000000O,
         oKXXXXXXXXXXXXO;;;:oxO0000000O.
          'kXXXXXXXXXX0:;;;;;;;cdkO000O.
            cKXXXXXXXK:;;;;;;;;;;;:ldOO
             .xXXXXXKl;;;;;;;;;;;;;,'..
               ;0XXXo;;;;;;;,'...
                 oXx;;,...
                  .
`,

	`                      ....',;;.
                .ckdkOO00000000o.
            .';oKMMkxxxxxkkkOO000x;
          'cllkWMMMNxxxxxxxxxxxxkkOOl.
         .lloKMMMMMMOxxxxxxxxxxxxxxxxko
         clxNMMMMMMMWxxxxxxxxxxxxxxkO0O.
        :oKMMMMMMMMMMKxxxxxxxxxxxkO000k;
       ,kWMMMMMMMMMMMMkxxxxxxxkO000000k;.
      .KMMMMMMMMMMMMMMNxxxxxkO00000000k;.
      xNNWWWWWWMMMMMMMM0xxO00000000000k;.
       ;0XXXXXXXXXXXXNN0ok000000000000k;
         dXXXXXXXXXXXXK:;;:lxO00000000x.
          ,OXXXXXXXXXKl;;;;;;;coxO0000x
            lKXXXXXXXd;;;;;;;;;;;;cdk0d
             .OXXXXXO;;;;;;;;;;;;;;,''.
               cKXX0:;;;;;;;,'...
                .xKc;,'...

`,

	`                      ....',;:.
                .dxxkO000000000d'
            ..;dNMNxxxxxkkkOO0000x;
          .:ll0MMMMOxxxxxxxxxxxkkOOO:
          cldNMMMMMNxxxxxxxxxxxxxxxxxk;
         :lOWMMMMMMM0xxxxxxxxxxxxxxkO0d.
        :dNMMMMMMMMMWkxxxxxxxxxxkO0000d;
       ;0MMMMMMMMMMMMKxxxxxxxxk0000000d;'
      :NMMMMMMMMMMMMMMkxxxxkO000000000d;,
      0NNWWWWWWMMMMMMMXxxk000000000000o;.
       cKXXXXXXXXXXXNNXok0000000000000o;
        .xXXXXXXXXXXXXo;;:ldk000000000o.
          ;0XXXXXXXXXk;;;;;;;:lxO00000o
            dXXXXXXX0:;;;;;;;;;;;:oxO0:
             ,0XXXXKl;;;;;;;;;;;;;;,',.
               oKXXd;;;;;;;;,'...
                'OO;;,'...

`,

	`                      ....',;:'
                'kdxOO000000000x,
             .;kWM0xxxxkkkOO00000x,
           ;loXMMMWxxxxxxxxxxxkkOOOk;
          :lkWMMMMM0xxxxxxxxxxxxxxxxkx.
         :oKMMMMMMMWkxxxxxxxxxxxxxkO0Kc
        ;kWMMMMMMMMM0xxxxxxxxxxxk0KKKKc;
       cKMMMMMMMMMMMWkxxxxxxxkO0KKKKKKc;,
      dWMMMMMMMMMMMMMKxxxxkO0KKKKKKKKKc;;
     .KNNNWWWWWMMMMMMWkxk0KKKKKKKKKKKKc;'
       lKKKKKKKKXXXXXNdx0KKKKKKKKKKKKK:;
        .OKKKKKKKKKKKk;;;cok0KKKKKKKKK:.
          cKKKKKKKKKK:;;;;;;:ldk0KKKKK:
           .kKKKKKKKo;;;;;;;;;;;:ldOKK.
             cKKKKKO;;;;;;;;;;;;;;;,,;
              .xKKKc;;;;;;;;,'...
                ;Kd;;,'..

`,

	`                      ....',;:,
                :kdkO0000000000k,
             .:0MWkxxxxkkOOO00000x,
           'cxNMMMKxxxxxxxxxxkkkOO0x'
          ,l0MMMMMWkxxxxxxxxxxxxxxxkko
         ,dNMMMMMMM0kxxxxxxxxxxxxxk0K0,
        :0MMMMMMMMMWkxkxxxxxxxxkOKKKK0;,
       oNMMMMMMMMMMM0xkxxxxxkO0KKKKKK0;;,
      OMMMMMMMMMMMMMNkkxxkO0KKKKKKKKKO;;;.
     'KNNNNWWWWMMMMMMOkkOKKKKKKKKKKKKO;;,
      .oKKKKKKKKXXXXXOxOKKKKKKKKKKKKKO;;
        ,0KKKKKKKKKK0:;;cox0KKKKKKKKKO;.
          oKKKKKKKKKd;;;;;;;cox0KKKKKO,
           'OKKKKKKO;;;;;;;;;;;;cdk0Kk.
             oKKKKKl;;;;;;;;;;;;;;;,:;
              'OKKk;;;;;;;;;,'...
                lK:;,'...

`,

	`                      ....',;;;
                lkxkO0000000000O;
             .cXMNkkkkkkOOO000000x'
           .cOWMMMOkkkkkkkkkkkkOO00d.
          .oXMMMMMKkkkkkkkkkkkkkkxkkk;
         'kWMMMMMMWkkkkkkkkkkkkkkkO0Kx'
        :XMMMMMMMMM0kkkkkkkkkkkO0KKKKx;,
       dWMMMMMMMMMMNkkkkkkkkO0KKKKKKKx;;;
     .KMMMMMMMMMMMMMOkkkkk0KKKKKKKKKKd;;;'
     ;XXNNNWWWWMMMMMXkk0KKKKKKKKKKKKKd;;;
      .dKKKKKKKKKXXXKdOKKKKKKKKKKKKKKd;;.
        ;0KKKKKKKKKKo;;:lxOKKKKKKKKKKd;.
         .dKKKKKKKKO:;;;;;;:oxOKKKKKKo,
           ,0KKKKKKl;;;;;;;;;;;:ox0KKo
            .xKKKKO;;;;;;;;;;;;;;;;:c.
              ;0KKl;;;;;;;;,'....
               .dk;;,'...

`,

	`                     .....',;;:.
               .xxkO00KKKKKKKKK0:
             .oNM0kkkkkOOO00KKKKKx.
            :0MMMNkkkkkkkkkkkkOO000l
          .oNMMMMMkkkkkkkkkkkkkkkkkOk.
         '0MMMMMMMKkkkkkkkkkkkkkkO0KXc.
        cNMMMMMMMMNkkkkkkkkkkkk0KXXXXc:,
      .kMMMMMMMMMMMOkkkkkkkk0KXXXXXXXc::;.
     ,NMMMMMMMMMMMMXkkkkkOKXXXXXXXXXXc:::;
     cXXNNNWWWMMMMMWkkOKKXXXXXXXXXXXXc::;
      .xKKKKKKKKKXXXdkKXXXXXXXXXXXXXX:::.
        :0KKKKKKKKKk;;:ldkKXXXXXXXXXX::.
         .xKKKKKKKKl;;;;;;:ldkKXXXXXX:'
           :0KKKKKO;;;;;;;;;;;:ldOKXX,
            .kKKKKl;;;;;;;;;;;;;;;:co
              c0KO:;;;;;;;;,'...
               .ko;;,'..

`,

	`                     .....',,;:.
               .kxk00KKKKKKKKKKKc
             .dWMkkkkkkOO000KKKKKx.
            ;XMMM0kkkkkkkkkkkOO00K0;
           dWMMMMXkkkkkkkkkkkkkkkkkOo
         'KMMMMMMWkkkkkkkkkkkkkkkOKX0;.
        lWMMMMMMMM0kkkkkkkkkkkOKXXXX0::,
      .0MMMMMMMMMMXkkkkkkkkO0XXXXXXX0:::;.
     :WMMMMMMMMMMMWkkkkkO0XXXXXXXXXX0:::::
     cKXXNNNWWMMMMM0kOKXXXXXXXXXXXXXO::::.
      .kKKKKKKKKKXXkk0XXXXXXXXXXXXXXO:::.
        c0KKKKKKKK0:;:cdk0XXXXXXXXXXO::.
         .kKKKKKKKx;;;;;;:cdk0XXXXXXO:.
           cKKKKKKl;;;;;;;;;;;cdkKXXk'
            'OKKKO:;;;;;;;;;;;;;;:coc
              oKKd;;;;;;;;;,''...
               'O:;,'...

`,

	`                     .....',,;:,
               ;kkO00KKKKKKKKKKKl
             .kMNkkkkkOO000KKKKKKx.
            ;NMMWkkkkkkkkkkkOO000K0'
           dMMMMM0kkkkkkkkkkkkkkkkOO;
         'XMMMMMMKkkkkkkkkkkkkkkO0KXx;.
        oMMMMMMMMNkkkkkkkkkkkO0KXXXXdcc,
      .KMMMMMMMMMMOkkkkkkkO0KXXXXXXXdccc:.
     lWMMMMMMMMMMM0kkkkO0KXXXXXXXXXXdccccc.
     cKXXNNNWWMMMMXkO0XXXXXXXXXXXXXXdcccc.
      .k000000KKKXOx0XXXXXXXXXXXXXXXdccc.
        l000000000o::cok0XXXXXXXXXXXocc.
         'k0000000::::::::oxOXXXXXXXoc.
           o00000k::::::::::::ox0XXXo.
            ,O000o:::::::::::::::cox'
             .d00c::::::::;;,'...
               ;k:;,'..

`,

	`                     .....',,;;;
               lkkO0KKKKKKKKKKKKd
             .0MKkkkkOOO00KKKKKKKx.
            cWMMXkkkkkkkkkkOO000KKk.
          .OMMMMWkkkkkkkkkkkkkkkkOOk'
         ;NMMMMMMOkkkkkkkkkkkkkO0KXXc;.
        kMMMMMMMM0kkkkkkkkkkO0KXNNNXccc,
      ,XMMMMMMMMMKOOOOOOkO0KXNNNNNNXccccc.
     dMMMMMMMMMMMNOOOOO0KXNNNNNNNNNXcccccc.
     cKKXXNNWWMMMWOO0XXNNNNNNNNNNNNXccccc,
      .k0000000KKKdOXNNNNNNNNNNNNNNKcccc'
        l00000000k:::lxOXNXNNNNNNXNKccc.
         'O000000d:::::::ldOKNNNNNNKcc.
           d00000l:::::::::::ldOXNNKc.
            :0000::::::::::::::::lxx
             .x0k:::::::::;;,'...
               cd:;,'..

`,

	`                    ......'',;;:
               xkO0KKKKKKKKKKKKKx
             ,XMOkkkOOO00KKKKKKKKx.
            dMMM0kkkkkkkkkOOO00KKKd.
          .KMMMMKOOOOOOOkkkkkkkkOO0l'.
         lWMMMMMXOOOOOOOOOOOOOkO0XNOc;.
       .0MMMMMMMNOOOOOOOOOOOOKXNNNNOllc,
      :NMMMMMMMMWOOOOOOOO0KXNNNNNNNkllllc'
     kMMMMMMMMMMMOOOOO0KXNNNNNNNNNNkllllll,
     :0KXXNNWWMMM0O0KNNNNNNNNNNNNNNklllll;
      .k0000000KKxkKNNNNNNNNNNNNNNNkllll'
        l00000000c::cdOKNNNNNNNNNNNklll.
         ,O00000O:::::::cdkKNNNNNNNklc.
          .d0000k:::::::::::cdkKNNNx:
            c000d:::::::::::::::cdOc
             .k0o:::::::::;;,''..
               oc;,,'..

`,

	`                    ......'',,;:.
              .kkO0KKXXXXXXXXXXXk.
             cWWOOOOOO00KKKXXXXXXd.
           .OMMWOOOOOOOOOOOO00KKXKc.
          ,NMMMMOOOOOOOOOOOOOOOOO00,'.
         dMMMMMMOOOOOOOOOOOOOOO0KXNoc,.
       .KMMMMMMM0OOOOOOOOOOO0KNNNNNoooc,.
      lWMMMMMMMMKOOOOOOOO0XNNNNNNNNoooooc'
     OMMMMMMMMMMKOOOO0KXNNNNNNNNNNNooooooo:
     :0KKXNNWWMMXO0KNNNNNNNNNNNNNNNoooooo:
      .x0000000KOx0NNNNNNNNNNNNNNNNooooo,
        l0000000d:::okKNNNNNNNNNNNNoool.
         ,O00000o::::::cox0XNNNNNNNoo:
          .x0000l:::::::::::ox0NNNXo,
            l000c::::::::::::::cokO.
             'k0::::::::::;;,''..
              .d:;,''..

`,

	`                    ......'',,;;.
              ,OO00KXXXXXXXXXXXXO.
             dMXOOOOOO00KKXXXXXXXl.
           .KMMXOOOOOOOOOOO00KKXXK,'.
          :WMMMXOOOOOOOOOOOOOOOOO0d''.
         kMMMMMNOOOOOOOOOOOOOOOKXN0l:,.
       'NMMMMMMNOOOOOOOOOOO0KXNWNN0oooc,.
      oMMMMMMMMWOOOOOOOO0KNNNNNNNN0oooooc'
     0MMMMMMMMMWOOOOOKXNWNNNNNNNNN0oooooooc
     ;O0KXXNWMMWO0KXNWWNNWNWNNNNNN0ooooooc
      .xOOOOOO00dOXWWNNNNNNNNNNNNN0ooooo,
        lOOOOOOO:::lx0XNNNNNNNNNNNOoool.
         ,OOOOOk:::::::lxOXNNNNNNNOoo;
          .xOOOk:::::::::::ldOXNNNOo.
            lOOk:::::::::::::::lx0d.
             ;Ox::::::::::;;,,'..
              .o;;,'...

`,

	`                   .......'',,;;'
              cOO0KKXXXXXXXXXXXX0.
             kM0OOOOO00KKXXXXXXXXc'
           'NMM0OOOOOOOOOOO00KXXXk,'.
          lMMMM0OOOOOOOOOOOOOOOO0K;',.
        .OMMMMM0OOOOOOOOOOOOOO0KNWxl:,'
       ;NMMMMMM0OOOOOOOOOOO0XNWWWWxddo:,.
      dMMMMMMMM0OOOOOOO0KXWWWWWWWWdddddoc'
     0WMMMMMMMM0OOOO0XNWWWWWWWWWWWddddddddl
     ,O0KXXNWMM0OKXWWWWWWWWWWWWWWWdddddddl.
      .dOOOOOO0xkKWWWWWWWWWWWWWWWWdddddo,
        lOOOOOOl::cdOXWWWWWWWWWWWWddddl.
         ,kOOOOl::::::cdkKNWWWWWWWddo,
          .xOOOo::::::::::cdkKWWWWdl.
            oOOo::::::::::::::cdOX;
             ;Oo::::::::::;;;,'..
              .l;,,'...

`,

	`                   .......'',,,;,
              okO0KXXXXXXXXXXXXXK.
            .0WOOOOOO00KKXXXXXXXX:'.
           ;NWWOOOOOOOOOOO00KKXXXl,,.
          dWWWWOOOOOOOOOOOOOOOO00x,,,.
        .0WWWWWOOOOOOOOOOOOOOOKXWKoc;,'
       :NWWWWWNOOOOOOOOOOO0KNWWWWKdddl:,.
      xWWWWWWWNOOOOOOOOKXWWWWWWWWKdddddoc,
     0WWWWWWWWNOOOO0XNWWWWWWWWWWWKddddddddl
     .k00KXNWWNO0XWWWWWWWWWWWWWWWKdddddddo.
      .dOOOOOOkx0NWWWWWWWWWWWWWWW0dddddd,
        cOOOOOx:::okKWWWWWWWWWWWW0ddddl.
         ,kOOOk::::::cok0NWWWWWWW0ddo'
          .xOOk:::::::::::ox0NWWW0d:
            oOO:::::::::::::::okKO.
             :Oc::::::::::;;;,''.
              'c;,''...

`,

	`                   .......'',,,;;
              xO00KXXXXXXXXXXXXXK'
            'KXkkkkOO0KKXXXXXXXXX,,.
           cWWXOOOOOOOOOOO0KKXXXX;,,.
          xWWWKOOOOOOOOOOOOOOOO0K:,,,'
        .KWWWWKOOOOOOOOOOOOOO0XNMxoc,,,.
       :NWWWWW0OOOOOOOOOOOKNWMMMMxxxdl:,.
      xWWWWWWW0OOOOOOOKXNMMMMMMMMxxxxxxoc'
     ONWWWWWWW0OOO0KNWMMMMMMMMMMMxxxxxxxxxl
     .xO0KXNWWO0XNMMMMMMMMMMMMMMMxxxxxxxxo.
       okkkkkOdOXMMMMMMMMMMMMMMMMxxxxxxd;
        ckkkkkc::cx0NMMMMMMMMMMMMxxxxxl.
         ,kkkkl::::::lx0XWMMMMMWWxxxo'
          .xkko::::::::::cd0NWWWWxd;
            okd::::::::::::::lx0Nl.
             ck::::::::::;;;;,''.
              ,;,,'....

`,

	`                  ........''',,;;
             .kO0KKXXXXXXXXXXXXXX,
            ,XKkkkkOO0KKXXXXXXXX0,,.
           oWW0kkkkkkkkkOO0KXXXXO,,,.
          kWWWOkkkkkOOOOOOOOOO0Kk,,,,'
        'KWWWWOOOOOOOOOOOOOOOKNWXdl:,,,.
       cNWWWWNOOOOOOOOOOO0XWMMMMXxxxdc;,.
      xWWWWWWNOOOOOOO0KNWMMMMMMMKxxxxxxoc'
     kNWWWWWWXOOOOKXWMMMMMMMMMMMKxxxxxxxxxl.
     .dO0KXNWK0KNMMMMMMMMMMMMMMMKxxxxxxxxd'
       lkkkkkxkKWMMMMMMMMMMMMMMMKxxxxxxx:
        :kkkkd:::okXWMMMMMMMMMMMKxxxxxl.
         'xkkk::::::cdOXWMMMMMMMKxxxo'
          .xkkc:::::::::cokXWMMMKxd,
            dkl:::::::::::::cdON0:.
             cd::::::::::;;;;,,'.
              ,;,''....

`,

	`                  ........''',,,;.
             .kO0KXXXXXXXXXXXXXXX;.
            :N0kkkkO00KXXXXXXXXXO;;.
           dNNOkkkkkkkkkO00KXXXXo;;;'
         .ONNNkkkkkkkkkkkkkkOO0Xc;;;;,
        'KNNNXkkkkkkkkkkkkkO0XNMkoc;;;,.
       cNNNNNKkkkkkkkkkO0KNWMMMMkkkxoc;;.
      xNNNNNN0kkkkkOOKXWMMMMMMMMkkkkkkxoc'
     dXNNNNNNOOOO0XNMMMMMMMMMMMMkkkkkkkkkxl.
      lkO0KNNOKXWMMMMMMMMMMMMMMMkkkkkkkkkx,
       ckkkkkdOXMMMMMMMMMMMMMMMMkkkkkkkx:
        ;kkkkc::lx0NMMMMMMMMMMMWkkkkkxc.
         .xkko::::::lx0NMMMMMMMWkkkkl.
          .dkd::::::::::lx0WMMMWkkl'
            ok::::::::::::::okXWd,.
             ll::::::::::;;;;,,,.
              ,,,'.....

`,

	`                 .........''',,,;.
             'kO0KXXXXXXXXXXXXXXX;.
            cNkkkkkO0KKXXXXXXXXXx;;.
           xNXkkkkkkkkkOO0KXXXXX:;;;'
         .ONNKkkkkkkkkkkkkkkO0KO;;;;;,
        'KNNN0kkkkkkkkkkkkkOKNWKxo:;;;;.
       :NNNNNOkkkkkkkkkO0XWMMMMKkkkxo:;;.
      dNNNNNXkkkkkkO0XNMMMMMMMMKkkkkkkxo:'
     cXNNNNNKkkkOKNWMMMMMMMMMMMKkkkkkkkkkxl
      ckO0KN00XNMMMMMMMMMMMMMMMKkkkkkkkkkx;
       ;xxxxdkKWMMMMMMMMMMMMMMMKkkkkkkkx:
        'xxxd:::okXWMMMMMMMMMMMKkkkkkxc.
         .dxx::::::cdOXMMMMMMMMKkkkxc.
          .dxl:::::::::cdOXMMMMKkxc..
            od:::::::::::::lx0W0l'.
             lc:::::::::;;;;;,,,..
              ,,''....

`,

	`                ..........'''',,,.
             ;kO0KXNNNNNNNNNNNNNX:.
            lXkkkkkO0KXNNNNNNNNNo;;'
           kN0kkkkkkkkkO00KXNNNO;;;;,
         .ONNOkkkkkkkkkkkkkkO0Xc;;;;;,
        'KNNXkkkkkkkkkkkkkO0XNWkdl:;;;;.
       :XNNNKkkkkkkkkkk0KNMMMMWOOOkdl:;;.
      oNNNNN0kkkkkkOKNWMMMMMMMWOOOOOOkdl:'
     ,KXNNNNOkkO0XWMMMMMMMMMMMWOOOOOOOOOOxc
      ,xk0KXOKNMMMMMMMMMMMMMMMWOOOOOOOOOOk;
       'xxxxdOXMMMMMMMMMMMMMMMWOOOOOOOOx:
        .dxxc::cx0NMMMMMMMMMMMWOOOOOOx:.
         .dxo::::::okKWMMMMMMMWOOOOx:.
          .oxc:::::::::okKWMMMWOOx;..
            oo::::::::::::cdOXWx;..
             c::::::::::;;;;;,;'..
              ,''.....

`,

	`               ...........'''',,,.
             :kO0KNNNNNNNNNNNNNNXc'
            oKkkkkkO0KXNNNNNNNNXl;;'
           kXOkkkkkkkkkO0KXNNNNd;;;;,
         .OXXkkkkkkkkkkkkkkO0KO;;;;;;;
        .KXXKkkkkkkkkkkkkkOKNWKxoc;;;;;.
       ;XXXXOkkkkkkkkkOKXWMMMMKOOOkdc;;;.
      cXXXXXkkkkkkO0XNMMMMMMMMKOOOOOOkdl:'
     .0XXXX0kkk0KNMMMMMMMMMMMMKOOOOOOOOOOd:
      .dk0K0KXWMMMMMMMMMMMMMMMKOOOOOOOOOOk:
       .ddddx0WMMMMMMMMMMMMMMMKOOOOOOOOx:
        .odd:::okXWMMMMMMMMMMMKOOOOOOx;.
         .odc:::::cdOXMMMMMMMMKOOOOo;.
           od:::::::::lx0NMMMMKOOo,..
            lc::::::::::::okKW0o,..
             c:::::::::;;;;;,;;...
              '''.....

`,

	`               ............'''',,'
             ckO0XNNNNNNNNNNNNNNKc'
            d0kkkkkO0XNNNNNNNNNKc::,
           kXkkkkkkkkkkO0KNNNNXc::::;
         .OXKkkkkkkkkkkkkkkO0Xl::::::;
        .0XXOkkkkkkkkkkkkk0XNNOdl:::::;.
       'KXXKkkkkkkkkkk0XNMMMMN000Oxo::::.
      :KXXX0kkkkkkOKNWMMMMMMMN000000Okoc:.
      kKXXXkkkOKXWMMMMMMMMMMMN0000000000kd;
      .ok0K0XNMMMMMMMMMMMMMMMN00000000000k:
        odddOXWMMMMMMMMMMMMMMN000000000x:
         odl::cx0NMMMMMMMMMMMN000000Od;.
          ld::::::okKWMMMMMMMN0000Ol,.
           ll:::::::::dOXWMMMN00kc...
            l::::::::::::cx0NNx:....
             ::::::::::;;;;;;:'...
              ''......

`,

	`              .............''''','
             lkOKXNNNNNNNNNNNNNNKc,
            dOkkkkk0KXNNNNNNNNN0c::,
           kKkkkkkkkkkk0KXNNNNO:::::;
          kKOkkkkkkkkkkkkkO0KO:::::::;
        .OKKkkkkkkkkkkkkkOKNW0kdc:::::;.
       .0KK0kkkkkkkkkOKNWMMMMK000Odl::::.
      'KKKKkkkkkkO0XWMMMMMMMMK000000Oxoc:.
      oKKK0kkk0XWMMMMMMMMMMMMK0000000000ko,
       ck00KNWMMMMMMMMMMMMMMMK00000000000k:
        cdox0NMMMMMMMMMMMMMMMK000000000d:
         co:::okKWMMMMMMMMMMM0000000Oo,.
          cl:::::cdOXMMMMMMMM00000kc'.
           cc::::::::lxKWMMMM000d;...
            c::::::::::::oOXM0o,....
             ;::::::::;;;;;,:,....
              '.......

`,

	`              .............''''','
             lkOKXNNNNNNNNNNNNNNKl;
            dkxxxkO0KXNNNNNNNNNOc::;
           x0xxxxxxxxkO0XNNNNXd:::::;
          kKkxxxxxxkkkkkkkOKKl:::::::;
        .kK0xxxxxkkkkkkkO0XWNOxoc:::::;.
       .OKKkxxxxkkkkk0XWMMMMN0000koc::::.
      .OKK0xxxkkk0KNMMMMMMMMN0000000kdl::.
      ;0KKkxkOKNWMMMMMMMMMMMN00000000000xo.
       ;x00XWMMMMMMMMMMMMMMMN000000000000k;
        ;ookKWMMMMMMMMMMMMMMN0000000000d;
         :l::cdONMMMMMMMMMMMN0000000Ol'.
          :c:::::lkKWMMMMMMMN00000x:...
           ::::::::::dONMMMMN00Oo,...
            ::::::::::::lxKWNkc.....
             ;:::::::;;;;;,;:'....
              ........

`,

	`              .............'''''''
             okOKNNNNNNNNNNNNNNNKl;
            dxxxxxO0XNNNNNNNNNNxc::;
           xOxxxxxxxxkOKXNNNNKl:::::;
          x0xxxxxxxxxxxxxk0Xkc:::::::;
         xKkxxxxxxxxxxxk0KNW0kdl::::::;
        xK0xxxxxxxxkOKNMMMMMKKKKOxoc::::
       kKKkxxxxxOKXWMMMMMMMMKKKKKKK0kdl::.
      .OK0xxk0XWMMMMMMMMMMMMKKKKKKKKKKKOxl.
       .k0KNMMMMMMMMMMMMMMMMKKKKKKKKKKKKKk,
        .ldOXMMMMMMMMMMMMMMMKKKKKKKKKK0d;
         'c::lxKWMMMMMMMMMMWKKKKKKKKkc..
          ,:::::cdONMMMMMMMWKKKKK0o;...
           ,::::::::lkKWMMMWKKKkc....
            ;::::::::::cdONW0d;.....
             ,::::::;;;;;;;c,......
              .......

`,

	`              ..............''''''
             ok0KNNNNNNNNNNNNNNN0o:
            dxxxxkOKXNNNNNNNNNXxccc;
           dkxxxxxxxxk0XNNNNNkcccccc:
          dOxxxxxxxxxxxxkOKKocccccccc;
         o0kxxxxxxxxxxxO0XWXOxoccccccc;
        o0Oxxxxxxxxk0XWMMMMXKKK0kdlcccc;
       o00xxxxxk0XNMMMMMMMMXKKKKKKKOxocc;
       x0kxkOKNMMMMMMMMMMMMXKKKKKKKKKK0kdc
        x0XWMMMMMMMMMMMMMMMXKKKKKKKKKKKKKk.
        .lx0WMMMMMMMMMMMMMMXKKKKKKKKKKOo,
         .:::oONMMMMMMMMMMMXKKKKKKKKx:..
          .:::::lxKWMMMMMMMXKKKKKkl'...
           '::::::::d0NMMMMXKK0d;....
            ,::::::::::okXMXkc'.....
             ,::::::;;;;;,::'......
              .......

`,

	`              ...............'''''
             ok0XNNNNNNNNNNNNNNN0o:
            oxxxxkOKNNNNNNNNNNKdccc:
           oxxxxxxxxxOKNNNNNXdcccccc:
          lkxxxxxxxxxxxxO0Xkccccccccc;
         cOxxxxxxxxxxxk0KNW0kdlccccccc;
        :0kxxxxxxxxOKNMMMMWKKKKOxoccccc;
       ;0OxxxxxOKNWMMMMMMMWKKKKKKKKkdlcc;
       lOxxk0XWMMMMMMMMMMMWKKKKKKKKKKK0ko;
        xXWMMMMMMMMMMMMMMMWKKKKKKKKKKKKKKx.
         lkKWMMMMMMMMMMMMMWKKKKKKKKKKKOl'
         .;:cd0WMMMMMMMMMMWKKKKKKKK0d;..
          .:::::oOXMMMMMMMWKKKKKKxc....
           .:::::::lkXMMMMWKKKOl'....
            ':::::::::cxKWW0d;......
             ,:::::;;;;;,;c,.......
              ......

`,

	`              .................'''
             ok0XNNNNNNNNNNNNNNN0o:
            lxxxxk0XNNNNNNNNNN0occc:
           lxxxxxxxxk0XNNNNN0occcccc:
          :kxxxxxxxxxxxk0X0occccccccc;
         ,kxxxxxxxxxxxOKXWKOxocccccccc,
        'Oxxxxxxxxk0XWMMMMXXXXKkdlccccc,
       .Okxxxxk0XWMMMMMMMMXXXXXXXK0kolcc'
       :kxkOKNMMMMMMMMMMMMXXXXXXXXXXXKOxo.
       .0NWMMMMMMMMMMMMMMMXXXXXXXXXXXXXXKd
        .oOXMMMMMMMMMMMMMMXXXXXXXXXXXKkl'
         .;;lkXMMMMMMMMMMMXXXXXXXXXOo,..
          .;;;;cxKWMMMMMMMXXXXXX0d;....
           ';;;;;;:d0WMMMMXXXKx:.....
            ,;;;;;;;;:oONMXkl'......
             ,;;;;;;;;;,,c;'.......
              .....

`,

	`              ....................
             ok0XNNNNWNWWNWNWNNN0oc
            cxxxxk0XWWWWWWWWWNOoccc:
           :xxxxxxxxk0NWWWWXklcccccc;
          'xxxxxxxxxxxxOKXklccccccccc,
         .xxxxxxxxxxxk0XWN0kdlcccccccc'
        .xxxxxxxxxOKWMMMMNXXXX0xocccccc.
        dxxxxxOKNMMMMMMMMNXXXXXXXKOxoccc.
       ,xxOKNWMMMMMMMMMMMNXXXXXXXXXXX0kdl.
       ;KWMMMMMMMMMMMMMMMNXXXXXXXXXXXXXXKc
        ,d0NMMMMMMMMMMMMMNXXXXXXXXXXXKkc.
         ';:oONMMMMMMMMMMNXXXXXXXXXkl'..
          ';;;:okXMMMMMMMNXXXXXXOl,...
           ,;;;;;;lkXMMMMNXXX0o,.....
            ,;;;;;;;;lkXMNKd;.......
             ,;;;;;;;;;,:c''.......
              ....

`,

	`              ....................
             lk0XWWWWWWWWWWWWWWN0dc
            :xxxxkKNWWWWWWWWWXkolll:
           ,xxxxxxxxOKNWWWWKxlllllll;
          .dxxxxxxxxxxk0N0dllllllllll'
          dxxxxxxxxxxOKNWKOxolllllllll.
         oxxxxxxxOKNMMMMWXXXXKOdllllllc.
        cxxxxk0XWMMMMMMMWXXXXXXXX0kdlllc
       .xx0XWMMMMMMMMMMMWXXXXXXXXXXXX0xo:
       oNMMMMMMMMMMMMMMMWXXXXXXXXXXXXXXX0'
        :xKWMMMMMMMMMMMMWXXXXXXXXXXXXKxc.
         ,;:dKWMMMMMMMMMWXXXXXXXXXKxc'..
          ,;;;:d0WMMMMMMWXXXXXXKxc'...
           ,;;;;;:d0WMMMWXXXXkc'.....
            ,;;;;;;;:d0WWXkl'.......
             ,;;;;;;;;,;c;''.......
              ..

`,

	`              ....................
             lk0NWWWWWWWWWWWWWWNOdc
            :xxxxOKNWWWWWWWWWKkllll:
           'xxxxxxxx0NWWWWNOolllllll,
          .dxxxxxxxxxxOXKklllllllllll.
          oxxxxxxxxxk0XWX0kdllllllllll.
         lxxxxxxk0XWMMMMNNNNX0xollllllc
        :xxxxOKNMMMMMMMMNNNNNNNNKOxolll:
       'xOKNMMMMMMMMMMMMNNNNNNNNNNNNKOdl,
       OWMMMMMMMMMMMMMMMNNNNNNNNNNNNNNNXk.
       .lkXMMMMMMMMMMMMMNNNNNNNNNNNNN0d:.
         ;;ckXMMMMMMMMMMNNNNNNNNNN0d:..
          ,;;;lkXMMMMMMMNNNNNNX0d:'...
           ,;;;;;cxXMMMMNNNN0d;......
            ,;;;;;;;lkXMNKd:........
             ,;;;;;;;,,c:''''......
              .

`,

	`              ....................
             ;x0XWWWWWWWWWWWWWWNOo;
            'xxxxOKWWWWWWWWWWKxllll,
           .dxxxxxxk0NWWWWN0dlllllll.
           oxxxxxxxxxx0NNOollllllllll.
          lxxxxxxxxxxOKNKOxolllllllllc
         cxxxxxxxOKNMMMMXXXXKOxollllll:
        ;xxxxOKNWMMMMMMMXXXXXXXX0kxolll;
       'xOKNWMMMMMMMMMMMXXXXXXXXXXXX0kdl,
       kWMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXXx
        cxKWMMMMMMMMMMMMXXXXXXXXXXXXXOo;
         ,;cxKWMMMMMMMMMXXXXXXXXXXOo;..
          ,;;;cd0WMMMMMMXXXXXXXko;....
           ';;;;;:d0WMMMXXXXkl,......
            ';;;;;;;:d0WXOl,........
             .;;;;;;,,,;,'''''.....


`,

	`              ....................
             .dOXWWWWWWWWWWWWWWXkl.
            .dxxxkXWWWWWWWWWWKxllll.
            oxxxxxxkKWWWWWWKdllllllc.
           lxxxxxxxxxxKWW0dlllllllllc
          cxxxxxxxxxxk0X0kdllllllllll:
         ;xxxxxxxk0XWMMMXXXX0kdlllllll;
        ,xxxxO0XWMMMMMMMXXXXXXXX0kdllll,
       .xk0XWMMMMMMMMMMMXXXXXXXXXXXX0kdl'
       kWMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXXx
        :x0WMMMMMMMMMMMMXXXXXXXXXXXXXko,
         ,;:d0WMMMMMMMMMXXXXXXXXXXkl,..
          ';;;:oONMMMMMMXXXXXXKxc'....
           .;;;;;;lkNMMMXXXKx:'......
            .;;;;;;;;lkN0x:'........
             .;;;;;,,,,,''''''.....


`,

	`               ..................
              okXWWWWWWWWWWWWWWXxc
             lxxxkKWWWWWWWWWWKxlllc
            cxxxxxxkKWWWWWWKdllllllc
           :xxxxxxxxxkXWWXxlllllllll:
          ;xxxxxxxxxxxOXKxollllllllll;
         ,xxxxxxxxOKNMMMXXXKOxolllllll,
        'xxxxk0XNMMMMMMMXXXXXXXKOxdllll,
       .xk0XWMMMMMMMMMMMXXXXXXXXXXXX0kdl'
       xNMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXKd
        ;d0NMMMMMMMMMMMMXXXXXXXXXXXXKxl,
         ';;oONMMMMMMMMMXXXXXXXXXKxc'..
          .;;;;lkKMMMMMMXXXXXXOd:'....
           .;;;;;;cxKWMMXXXOo;.......
            .;;;;;;;;cd0kl,.........
              ,;;;;,,,,'''''''....


`,

	`               ..................
              :xKWWWWWWWWWWWWWWKd:
             ;xxxkKWWWWWWWWWWKdlll;
            ;xxxxxxkXWWWMWWXxllllll;
           ,xxxxxxxxxONWWNklllllllll,
          'xxxxxxxxxxxkXKdlllllllllll,
         'xxxxxxxxO0XWMMXXK0kdllllllll'
        .dxxxkOKNWMMMMMMXXXXXXX0Oxollll'
       .xk0XWMMMMMMMMMMMXXXXXXXXXXXKOkdl.
       dNMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXKd
        ,oOXMMMMMMMMMMMMXXXXXXXXXXXX0xc'
         .;;lkXMMMMMMMMMXXXXXXXXXOd:...
          .;;;;cx0WMMMMMXXXXXKkl;.....
            ;;;;;;:oONMMXXKxc'......
             ,;;;;;;;;lkdc'........
              ';;;;,,,''''''''....


`,

	`               ..................
              .x0WMMMMMMMMMMMMW0o'
             .dxxxKWMMMMMMMMMKdlll'
            .dxxxxxkXMMMMMMXxllllll'
           .dxxxxxxxxONMMWklllllllll'
          .dxxxxxxxxxxxXKolllllllllll.
         .dxxxxxxxkOKNWMXX0kxollllllll.
        .dxxxxOKXWMMMMMMXXXXXXKOkdollll.
       .dk0XNMMMMMMMMMMMXXXXXXXXXXXKOxdl.
       oNMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXKo
        'okXMMMMMMMMMMMMXXXXXXXXXXXXOd:.
         .;;cx0WMMMMMMMMXXXXXXXXKkl;...
          .;;;;:oOXMMMMMXXXXX0xc,....
            ,;;;;;;lxKWMXXOo;.......
             ';;;;;;;;:dl;.........
              .;;;;,,,''''''''...


`,

	`                .................
              .oOWMMMMMMMMMMMMWOl.
             .oxxxKWMMMMMMMMMKollc.
            .dxxxxxkXMMMMMMNxllllll.
           .dxxxxxxxx0WMMWOlllllllll.
          .dxxxxxxxxxxxXKolllllllllll.
         .dxxxxxxxxk0XNMXKOxdlllllllll.
        .dxxxxO0KNWMMMMMXXXXXK0Oxdollll.
       .dk0KNWMMMMMMMMMMXXXXXXXXXXX0Oxdl.
       oXMMMMMMMMMMMMMMMXXXXXXXXXXXXXXX0l
        .lxKWMMMMMMMMMMMXXXXXXXXXXXKko:.
         .;;:oONMMMMMMMMXXXXXXXX0xc,...
           ,;;;;lxKWMMMMXXXXKko:'....
            ';;;;;;:dONMXKxc,.......
             .;;;;;;;;;lc'.........
              .,;;;,,,,,,,,'''...


`,

	`                ................
               ckNMMMMMMMMMMMMNx:
              lxxx0WMMMMMMMMM0lllc
             oxxxxxkXMMMMMMNdlllllc.
            oxxxxxxxx0WMMM0llllllllc.
          .oxxxxxxxxxxkXXdllllllllllc.
         .dxxxxxxxxkOKXWKOkdolllllllll.
        .dxxxxk0KXWMMMMMXXXXXKOkxolllll.
       .dk0KNWMMMMMMMMMMXXXXXXXXXXK0kxdl.
       lKMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXOl
        .cx0NMMMMMMMMMMMXXXXXXXXXXXKxl;.
         .;;;lkXMMMMMMMMXXXXXXXKOd:'..
           ';;;;cdONMMMMXXXX0xl,.....
            .;;;;;;;lkKWKOo:'.......
             .;;;;;;;;;:;''........
               ';;;,,,,,,,,,''...


`,

	`                ................
               ,xXMMMMMMMMMMMMXo,
              ;xxx0WMMMMMMMMWOlll;
             cxxxxxkXMMMMMMNdlllll:
            cxxxxxxxx0MMMM0llllllll:
           lxxxxxxxxxxkNNxllllllllllc.
         .oxxxxxxxxxk0KN0kxolllllllllc.
        .dxxxxkO0XNWMMMMXXXXK0Oxdollllc.
       .dk0KXWMMMMMMMMMMXXXXXXXXXXK0kxol.
       cKMMMMMMMMMMMMMMMXXXXXXXXXXXXXXKOc
        .:dOXMMMMMMMMMMMXXXXXXXXXXX0xc,.
          ,;;cx0WMMMMMMMXXXXXXXKkl;...
           .;;;;:okKWMMMXXXKOd:'.....
            .;;;;;;;:dON0xc,........
              ,;;;;;;;;;,''.......
               .;;;;,,,,,,,,''...


`,

	`                 ...............
               .d0MMMMMMMMMMMMKl.
              'dxxOWMMMMMMMMWklll'
             ;xxxxxxXMMMMMMNdlllll;
            :xxxxxxxxKMMMMKllllllll;
           cxxxxxxxxxxOWWkllllllllll:
          oxxxxxxxxxxk0XOxdllllllllllc.
        .dxxxxxk0KXNMMMMKKKK0Okxdollllc.
       .xkOKXNMMMMMMMMMMKKKKKKKKKK0Okxol.
       :0WMMMMMMMMMMMMMMKKKKKKKKKKKKKKKk:
        .:okKWMMMMMMMMMMKKKKKKKKKKKOdc'.
          ,;;:dONMMMMMMMKKKKKKK0xc,...
           .;;;;;cx0NMMMKKK0xl;'.....
             ,;;;;;;;lxKko;'.......
              ';;;;;;;;;,,'.......
               .;;;;;,,,,,,,,'...
                 ............

`,

	`                 ..............
                lOWMMMMMMMMMMMkc.
              .oxxkNMMMMMMMMWxllc.
             .dxxxxxXMMMMMMNolllll'
            ,xxxxxxxxKMMMMKllllllll,
           :xxxxxxxxxxOWMOllllllllll;
          lxxxxxxxxxxxOKOdollllllllllc
        .oxxxxxkO0KNWMMMKKK0Okxdolllllc.
       .xkOKXNWMMMMMMMMMKKKKKKKKKK0Okdol.
       :0WMMMMMMMMMMMMMMKKKKKKKKKKKKKKKx:
        .;lxKWMMMMMMMMMMKKKKKKKKKKKko:'
          ';;:lkKWMMMMMMKKKKKKKko:'...
           .;;;;;:okKWMMKKKkdc,......
             ';;;;;;;:oOdc,........
              .;;;;;;;;;;,'.......
                ,;;;;;;;;;;,,'..
                 ............

`,

	`                 ..............
                ;kNMMMMMMMMMMWd;
               lxxxXMMMMMMMMNdll:
             .oxxxxxXMMMMMMXollllc.
            .dxxxxxxxKMMMMKllllllll'
           ;xxxxxxxxxx0MM0llllllllll,
          lxxxxxxxxxxxxK0olllllllllll:
        .oxxxxxxO0KXNWMMKK0Okkxdolllllc.
       .xkO0XNWMMMMMMMMMKKKKKKKKK0Okxdol.
       ;ONMMMMMMMMMMMMMMKKKKKKKKKKKKKK0x;
         ,lx0NMMMMMMMMMMKKKKKKKKKK0xl;.
          .;;;cx0NMMMMMMKKKKKK0xl;....
           .,;;;;;lx0NMMKK0xl;'.....
             .;;;;;;;;cxo;'........
              .;;;;;;;;;;;'.......
                ';;;;;;;;;;;,'..
                  ...........

`,

	`                 ..............
                .dXMMMMMMMMMMNo'
               ;xxxKMMMMMMMMXoll;
              lxxxxxKMMMMMMKlllllc.
            .dxxxxxxxKMMMMKlllllllc.
           ,dxxxxxxxxx0MMKllllllllll,
          cxxxxxxxxxxxxK0llllllllllll:
        .oxxxxxxkO0KXNWMK0OOkxdollllllc.
       .xkO0KXNWMMMMMMMMKKKKKKKKK0Okxdol.
       ,OXMMMMMMMMMMMMMMKKKKKKKKKKKKKKOd,
         ,cdOXWMMMMMMMMMKKKKKKKKKKOdc,.
          .;;;:okXWMMMMMKKKKKKOdc,....
            ';;;;;:okKWMK0koc,......
             .;;;;;;;;:oc'.........
               ,;;;;;;;;;;,'.....
                .;;;;;;;;;;;,'..
                  ...........

`,

	`                  ............
                .o0MMMMMMMMMM0c.
               'dxx0MMMMMMMMKlll'
              :xxxxx0MMMMMMKlllll;
            .oxxxxxxxKMMMMKlllllllc.
           'dxxxxxxxxxKMMKllllllllll'
          :xxxxxxxxxxxxKKllllllllllll;
        .oxxxxxxxkO0KXNW0Okkxdoollllllc.
       .xkO0KXNWMMMMMMMMKKKKKKKK0OOkxdol.
       'kXWMMMMMMMMMMMMMKKKKKKKKKKKKKKOo,
         ':okKWMMMMMMMMMKKKKKKKKK0ko:,.
          .;;;:lx0NMMMMMKKKKK0xl;'....
            ';;;;;;lxOXMKOdl;'......
             .,;;;;;;;;c:'.........
               ';;;;;;;;;;,'.....
                .,;;;;;;;;;;,'..
                  ...........

`,

	`                  ............
                 :kWMMMMMMMMMk:
               .oxxOMMMMMMMMOllc.
              ,dxxxxOMMMMMM0lllll,
             cxxxxxxx0MMMMKlllllll:
           .dxxxxxxxxxKMMXlllllllllc.
          ;xxxxxxxxxxxxXXllllllllllll,
        .lxxxxxxxkkO0KXNOkkxdoolllllll:.
       .xkO0KXXNWMMMMMMMKKKKKKK00Okxxdol.
       'xKWMMMMMMMMMMMMMKKKKKKKKKKKKK0ko'
         .:lx0XMMMMMMMMMKKKKKKKKKOxl:'.
          .,;;;cdOXWMMMMKKKKKkdc,....
            .;;;;;;:ok0NOxo:,.......
              ,;;;;;;;;:;'........
               .;;;;;;:::;,'.....
                 ';;;::::::;,..
                  ..........

`,

	`                  ............
                 ,xNMMMMMMMMWo,
                cxxkWMMMMMMWxll:.
              'dxxxxOMMMMMMklllll'
             :xxxxxxx0MMMM0lllllll;
           .oxxxxxxxxxKMMXlllllllllc.
          ;xxxxxxxxxxxxXNolllllllllll,
        .lxxxxxxxxkkO0KXOxxdoollllllll:.
       .xkO00KXNWMMMMMMMK0K00K00Okkxddol.
       .d0NMMMMMMMMMMMMMK0K00K00K00KK0xl.
         .;cdOKWMMMMMMMMK0K00K000kdc;'.
           ,;;;:lx0NMMMMK0K0Oxl:'....
            .;;;;;;;cdOKkoc;'.......
              ';;;;;;;;:;'........
               .;;;;;;::::,'.....
                 .;;:::::::;'..
                   .........

`,

	`                   ..........
                 .oXMMMMMMMMXl.
                ;dxxNMMMMMMWoll;
              .oxxxxkWMMMMMxllllc.
             ,dxxxxxxOMMMM0lllllll,
           .lxxxxxxxxxKMMXlllllllll:.
          ,dxxxxxxxxxxxNNolllllllllll'
        .lxxxxxxxxxkkO0Kkddoolllllllll:.
       .xkOO0KXXNWMMMMMM0000000OOkxxddol.
       .d0XMMMMMMMMMMMMM0000000000000Odc.
         .;cok0NMMMMMMMM00000000Oxo:,..
           ';;;;cdOKWMMM0000koc,.....
            .,;;;;;;:lx0dl:'........
              .;;;;;;;;::,........
                ,;;;;:::::,.....
                 .;;:::::::;'..
                   .........

`,

	`                   ..........
                  cOMMMMMMMM0:.
                'dxxXMMMMMMNlll'
               lxxxxxWMMMMWdllll:.
             'dxxxxxxOMMMMOlllllll'
            cxxxxxxxxxKMMXlllllllll:
          'dxxxxxxxxxxxNWolllllllllll'
         cxxxxxxxxxxxkOKkooollllllllll:.
       .xkOO00KXXNWWMMMM00000OOkkxxddool.
       .oOXWMMMMMMMMMMMM0000000000000kdc.
         .;:lxOXWMMMMMMM00000000kdl;'..
           .;;;;:lx0NMMM000Odl:'.....
             ,;;;;;;;coko:,........
              .;;;;;;;:::,........
                ';;;;:::::;.....
                  ,;:::::::;'.
                   .........

`,

	`                   ..........
                  ;xWMMMMMMMx,
                .oxx0MMMMMM0llc.
               :xxxxxNMMMMWollll;
             .oxxxxxxkMMMMkllllllc.
            :xxxxxxxxxKMMXlllllllll;
          .dxxxxxxxxxxxNWollllllllllc.
         cxxxxxxxxxxxxk0kollllllllllll;
       .xkkO00KKXXNNWMMM0000OOkkxxxddool.
       .lkKWMMMMMMMMMMMM000000000000Oxo:.
         .,;ldkKNMMMMMMM0000000Oxoc;'..
           .;;;;;cdkKNMM00Oxoc,......
             ';;;;;;;;cdl;'........
              .,;;;;;;:::,.......
                .;;;;:::::,.....
                  .;:::::::;..
                    .......

`,

	`                    .........
                  .dNMMMMMMWl'
                .cxxkMMMMMMkll:.
               ,dxxxxXMMMMNlllll,
             .lxxxxxxxWMMMxllllll:.
            ;dxxxxxxxx0MMKlllllllll,
          .oxxxxxxxxxxxWWollllllllllc.
         :xxxxxxxxxxxxxOklllllllllllll;
       .xkkOO00KKXXNNWWW0OOOkkkxxxdddool.
       .ck0NMMMMMMMMMMMM000000000000Odl;.
         .,;coxOXWMMMMMM0000000kdl:,..
           .;;;;;:lxOXWM00koc;'......
             .;;;;;;;;:l:'.........
               ';;;;;;:cc,.......
                .,;;;:cccc,.....
                  .;:cccccc,..
                    .......

`,

	`                    ........
                  .lKMMMMMMXc.
                 ;dxxWMMMMMdll,
               .oxxxx0MMMMKllllc.
              cxxxxxxxWMMMollllll:
            ,dxxxxxxxx0MM0llllllllc'
          .oxxxxxxxxxxxWWollllllllllc.
         :xxxxxxxxxxxxx0Olllllllllllll;
       .xkkOOO00KKKXXXNNOkkkkxxxddddoool.
        :xOXWMMMMMMMMMMM000000000000kdl;.
          ';:ldkKNMMMMMM000000kxoc;'..
           .,;;;;;cox0XW0kdl:,.......
             .;;;;;;;;;c;..........
               .;;;;;;:cc,.......
                 ';;;:cccc,....
                  .,;cccccc'..
                    .......

`,

	`                    ........
                   :kMMMMMMk;
                 'dxxNMMMMWllc'
               .lxxxxOMMMMOllllc.
              ;dxxxxxxNMMWollllll;
            'oxxxxxxxxOMM0llllllllc.
          .lxxxxxxxxxxxWWollllllllll:.
         :xxxxxxxxxxxxx0Olllllllllllll,
       .xkkkOOO000KKKKXXkkxxxxdddddooool.
        ;dOKNMMMMMMMMMMM00000000000Oxoc,
          ';;coxOKWMMMMM00000Oxol:,'..
           .,;;;;;;cdk0Xkdl:,'......
             .,;;;;;;;;:;..........
               .;;;;;;:cc,.......
                 .;;;;cccc'....
                   ';cccccc'.
                     .....

`,

	`                     .......
                   'dWMMMMMo'
                 .lxxKMMMMXllc.
                cxxxxxMMMMxllll:.
              ,dxxxxxxXMMNlllllll'
            .oxxxxxxxxOMMOllllllllc.
          .cxxxxxxxxxxxNWollllllllll:.
         ;dxxxxxxxxxxxx00lllllllllllll,
       .xkkkOOOOO00000KKxxdddddddooooool.
        ,dk0XWMMMMMMMMMM0000000000Okdl:'
          .;;:ldk0XWMMMM0000Okdlc;'...
            ';;;;;;:ldk0dlc;'.......
             .,;;;;;;;;:;.........
               .,;;;;;:cc,.......
                 .;;;;cccc'....
                   .;:cccc:..
                     .....

`,

	`                     ......
                   .lXMMMMNc.
                 .cxxOMMMMOll:.
                ;dxxxxWMMMollll,
              'oxxxxxxKMMXllllllc.
            .lxxxxxxxxkMMkllllllll:.
           cxxxxxxxxxxxNWlllllllllll;
         ;dxxxxxxxxxxxx00lllllllllllll,
       .xkkkkkOOOOOOOOO0xdddoooooooooool.
        'oxOKWMMMMMMMMMMOOOOOOOOOOOxol:.
          .;;;coxOKNWMMMOOOOkdoc:,'...
            ';;;;;;;:ldkoc;'........
              ';;;;;;;;::.........
                ,;;;;;:cl,......
                 .,;;;cllc.....
                   .,:llll;..
                     .....

`,

	`                     ......
                    :0MMMM0:.
                  ;dxxMMMMdll,
                ,dxxxxXMMWllllc'
              .oxxxxxx0MM0llllllc.
            .cxxxxxxxxxMMxllllllll:.
           :xxxxxxxxxxxNWlllllllllll;
         ,dxxxxxxxxxxxx00llllllllllllc'
       .xkkkkkkkkkkkkkkOdooooooooooooool.
        .ldk0XWMMMMMMMMMOOOOOOOOOOkdoc;.
          .;;;:ldxOKNWMMOOOkxol:;'....
            .;;;;;;;;:ldc;,.........
              .;;;;;;;;c:.........
                ';;;;;;ll,......
                  ';;;cll:....
                   .':llll,.
                     .....

`,

	`                      .....
                    ,xMMMMx,
                  'oxxNMMWllc'
                .oxxxxKMMXllllc.
              .cxxxxxxOMMOllllll:.
            .:dxxxxxxxxWMdllllllll;.
           ;dxxxxxxxxxxXNlllllllllll,
         ,dxxxxxxxxxxxx0Ollllllllllllc'
       .dkkkkkkkxxxxxxxkdllllllllllooool.
        .cdkOKNMMMMMWMWWOOOOOOOOOkxol:;.
          .,;;;cldkOKNWMOOkxolc;,'....
            .;;;;;;;;;:o:,..........
              .;;;;;;;;c:.........
                .;;;;;;ll'......
                  .;;;:ll:....
                    .;clll'.
                      ...

`,

	`                      ....
                    .oNMMWc.
                  .lxxKMMXll:.
                .cxxxxOMMOllll:.
               :xxxxxxxMMxllllll;.
             ;dxxxxxxxxWMlllllllll,
           ,dxxxxxxxxxxXXlllllllllll,
         ,oxxxxxxxxxxxx0Ollllllllllllc'
       .dkkkxxxxxxxxxxxkdllllllllllllooc.
        .:oxO0XNWWWWWNNNkkkkOOOOkxdlc:,.
          .,;;;:codk0KNWOkxolc;,'.....
            .,;;;;;;;;;c,...........
              .,;;;;;;;c:.........
                .,;;;;;ll'......
                  .;;;:ll;....
                    .;cllc..
                      ...

`,

	`                      ....
                    .cKMMK:.
                  .:dxOMMOll;.
                 :dxxxxMMxllll;.
               ;dxxxxxxWMollllll,
             ,dxxxxxxxxNWllllllllc'
           ,dxxxxxxxxxxKKllllllllllc'
         'oxxxxxxxxxxxxOOllllllllllllc.
       .okkxxxxxxxxxxxxkdllllllllllllloc.
        .;ldkO0XNNNNNXXXxxxkkkkkxdol:;'.
          .,;;;;:codkOKNxdolc:,'.....
            .,;;;;;;;;;:'..........
              .,;;;;;;;c:........
                .,;;;;;cl......
                  .,;;;ll,....
                    .,:ll;..
                      ...

`,

	`                      ....
                     ;kMMk,
                   ;dxxMMdll,
                 ,dxxxxWMollll,
               'dxxxxxxNWllllllc'
             'oxxxxxxxxXXllllllllc'
           'oxxxxxxxxxx00llllllllllc.
         'oxxxxxxxxxxxxOkllllllllllllc.
        okkxxxxxxxxxxxxkdllllllllllllllc.
        .,lodkO0KXXXXKKKxxxxxxxxdolc:,'
          .';;;;;:lodkOKdolc:,'......
             ';;;;;;;;;:'..........
               ';;;;;;;:;........
                 ';;;;;cc......
                   ';;;ll'...
                     ';ll,.
                       .

`,

	`                       ..
                     'oWMl'
                   'oxxNMllc'
                 'oxxxxXNllllc'
               .oxxxxxxKXllllllc.
             .lxxxxxxxx00llllllllc.
           .oxxxxxxxxxxOOllllllllllc.
         .lxxxxxxxxxxxxOxllllllllllll:.
        lxkxxxxxxxxxxxxkdllllllllllllll:.
         'clodxkO0KKK000ddddxxdolcc:;,.
           ';;;;;;:lodxOlc:;,'.......
             ';;;;;;;;;:'..........
               .;;;;;;;:;........
                 .;;;;;c:......
                   .;;;cc....
                     .;ll..
                       .

`,

	`                       ..
                     .lXX:.
                   .cxxKKll:.
                 .lxxxx00llll:.
               .lxxxxxx0Ollllll:.
             .cxxxxxxxxOkllllllll:.
           .lxxxxxxxxxxOxllllllllll:.
         .lxxxxxxxxxxxxkdllllllllllll:.
        cxxxxxxxxxxxxxxxollllllllllllll:
         ':clodxxkO00OOOoooddoolc::;,'.
           .;;;;;;;:codxc:;,'........
             .;;;;;;;;;:'..........
               .;;;;;;;:,........
                 .;;;;;:;......
                   .;;;c:....
                     .,cc..
                       .

`,

	`                       ..
                     .:Ok;.
                   .:dxOkll;.
                 .:xxxxOxllll;.
               .:xxxxxxkxllllll;.
             .:dxxxxxxxkdllllllll;.
           .cxxxxxxxxxxkdllllllllll:.
         .cxxxxxxxxxxxxxollllllllllll:.
        cdxxxxxxxxxxxxxxllllllllllllllc;
         .;:clloddxkkkkxlloollc::;;,''.
           .;;;;;;;;:clo;;,'.........
             .;;;;;;;;;;'..........
               .,;;;;;;:'........
                 .,;;;;:,......
                   .,;;:,...
                     .':;


`,

	`
                      ,oc'
                    ,dxxllc,
                  ;dxxxxlllll,
                ;dxxxxxxlllllll,
             .;dxxxxxxxxlllllllll;.
           .:xxxxxxxxxxxlllllllllll;.
         .cxxxxxxxxxxxxxlllllllllllll:.
        :dddxxxxxxxxxxxxlllllllllllllcc;
         .;:::ccllooddxxllcc:::;;,,''..
           .,;;;;;;;;;::''...........
             .,;;;;;;;;;...........
               .';;;;;;;........
                  ';;;;;......
                    ';;;....
                      .;.


`,

	`
                      ,oc'
                    ;dxdlll,
                  ;dxxxdlllll,
                ;dxxxxxdlllllll,
             .:dxxxxxxxdlllllllll;.
           .:dxxxxxxxxxdlllllllllll;.
         .cxxxxxxxxxxxxolllllllllllll;.
        :oddxxxxxxxxxxxollllllllllllllc;
         .,;::ccllooddxollccc::;;,,,''.
           .,;;;;;;;;;:;''''.'.'.'...
             .,;;;;;;;;,.''''''''..
               .';;;;;;,.''''''..
                  ';;;;,'..''.
                    ';;;..'.
                      .;.


`,

	`
                      ,ol,
                    ,dddlll,
                  ;dddddlllll,
                ;ddddddolllllll,
             .:ddddddddolllllllll;.
           .cddddddddddolllllllllll;.
         .cddddddddddddllllllllllllll;.
        ;ooddddddddddddllllllllllllllll;
         .,;::cclloodddlllccc::;;;,,''.
           .,;;;;;;;;;;,''''''''''''.
             .';;;;;;;;'''''''''''.
               .';;;;;;''''''''..
                  ';;;;,'''''.
                    ';;,'''.
                      .,'


`,

	`
                      ,ol,
                    ,oddlll,
                  ;ddddolllll,
                ;ddddddolllllll,
             .;ddddddddolllllllll,.
           .:ddddddddddllllllllllll;.
         .cddddddddddddllllllllllllll;.
        ,oodddddddddddollllllllllllllll;
         .';;::cllooddolllccc:::;;,,,'.
           .';;;;;;;;;,'''''''''''''.
              ';;;;;;;,'''''''''''.
                .;;;;;;''''''''..
                  .;;;;''''''..
                    .;;,'''.
                      .,'


`,

	`
                      ,ol,
                    ,odoool,
                  ;ddddoooool,
                ;ddddddoooooool,
             .;ddddddddoooooooool,
           .:dddddddddooooooooooool;
         .cdddddddddddooooooooooooool;.
        'looddddddddddooooooooooooooool;
          ';;::cllodddoolllccc:::;;,,,..
            .;;;;;;;;;,,''''''''''''.
              .;;;;;;;''''''''''''.
                .;;;;;,'''''''''.
                  .;;;;''''''..
                    .;;''''..
                      .,'


`,

	`
                      'ol,
                    ,odoool,
                  ;ddddoooool,
                ;ddddddoooooool,
             .;dddddddooooooooool,
           .:dddddddddooooooooooool,
         .:dddddddddddooooooooooooool,
        .loodddddddddoooooooooooooooooo,
          .,;::clloddoollllccc:::;;;,,'.
            .,;;;;;;;,,'''''''''''''.
              .,;;;;;,''''''''''''.
                .,;;;,''''''''''.
                  .,;;,'''''''.
                    .,,''''..
                      .,'


`,

	`
                      'ol,
                    ,odoool,
                  ;odddoooooo,
                ,oddddooooooool,
             .;oddddddooooooooool,
           .:dddddddddooooooooooooo,
         .:ddddddddddoooooooooooooool,
        .clodddddddddoooooooooooooooooo,
          .,;;:cllodoooolllcccc:::;;;,'.
            .,,,,,,;;,,,,,,,,,,,,,,,.
              .,,,,,,,,,,,,,,,,,,'.
                .,,,,,,,,,,,,,,'.
                  .,,,,,,,,,,'.
                    .,,,',,'.
                      .,'


`,

	`
                      'lo,
                    ,odoool,
                  ,odddoooooo,
                ,oddddooooooool'
              ;oddddddooooooooool'
            ;odddddddoooooooooooool'
         .;odddddddddoooooooooooooool'
         :loodddddddooooooooooooooooool'
          .,,;:cllodoooollllccc::::;;;,.
            .',,,,,;,,,,,,,,,,,,,,,,.
              .,,,,,,,,,,,,,,,,,,,.
                .,,,,,,,,,,,,,,'.
                  .,,,,,,,,,,'.
                    .,,,,,,'.
                      .,,


`,

	`
                      'lo,
                    ,odoool,
                  ,odddoooooo'
                ,oddddooooooool'
              ,odddddoooooooooool'
            ;odddddddoooooooooooool.
         .;oddddddddooooooooooooooool.
         ;clodddddddooooooooooooooooool.
           ',;:clooooooollllcccc::::;;,.
            .',,,,,;,,,,,,,,,,,,,,,,..
              .',,,,,,,,,,,,,,,,,,.
                .',,,,,,,,,,,,,'.
                  .',,,,,,,,,'.
                    .,,,,,,'.
                      .,,


`,

	`
                      'lo,
                    'loddoo'
                  ,oododddodo'
                'ooooodddddoodl.
              ,ooooooddddddddoooc.
            ,ooooooooddddddooodoodc.
          ,ooooooooodddddooododoodooc.
         'cloooooooodddoddoooooddooodoc.
           .,;:cloodoooolllllcccc::::;,.
             .,,,,,;,,,,,,,,,,,,,,,,..
               ',,,,,,,,,,,,,,,,,,.
                .',,,,,,,,,,,,,,.
                  .',,,,,,,,,'.
                    .',,,,,'.
                      .,,


`,

	`
                      'lo,
                    'lodddo'
                  'looddddddo'
                'looooddddddddl.
              'looooodddddddddddc.
            'looooooddddddddddddddc.
          ,loooooooodddddddddddddddd:.
         .:loooooooddddddddddddddddddd:
           .,;:cloddooooollllccccc::::,.
             .,,,,,;,,,,,,,,,,,,,,,,.
               .,,,,,,,,,,,,,,,,,,.
                 .,,,,,,,,,,,,,,.
                   ',,,,,,,,,'.
                    .',,,,,'.
                      .,,


`,

	`
                      'lo,
                    'lodddo'
                  'looddddddl.
                .looodddddddddc.
              'loooooddddddddddd:.
            'loooooodddddddddddddd:
          'loooooooddddddddddddddddo;
         .;clooooooddddddddddddddddddo,
           .';:clodoooooolllllccccc:::;.
             .''',;,,,,,,,,,,,,,,,,,.
               .'',,,,,,,,,,,,,,,,.
                 .',,,,,,,,,,,,,.
                   .,,,,,,,,,'.
                    .',,,,,'.
                      .,,


`,

	`
                      .lo,
                    .lodddo'
                  .looddddddl.
                .cooodddddddddc.
              .coooodddddddddddd:
            .coooooodddddddddddddd;
          .coooooooddddddddddddddddo,
          'clooooodddddddddddddddddddo'
            .,:cldddoooooolllllcccccc:;
             .''',;;;;;;;;;;;;;;;;;,.
               .'',,;,;,;;,;,;,;,,.
                 .',,,;,;;,;,;,,.
                   .,,;,;;,;,,.
                     ',,;;,'.
                      .,,


`,

	`
                      .lo,
                    .codddl.
                  .cooddddddl.
                .coooddddddddd:
              .:oooodddddddddddd;
            .:oooooddddddddddddddo,
          .:oooooodddddddddddddddddl.
          .:looooddddddddddddddddddddl.
            .,:codddoooooollllllcccccc;
              .',;;;;;;;;;;;;;;;;;;,.
               .'',;;;;;;;;;;;;;;,.
                 .',;;;;;;;;;;;,.
                   .,;;;;;;;;'.
                     ';;;;;'.
                      .,;


`,

	`
                      .oo'
                    .cddddl.
                  .cloddddddc.
                .:lloddddddddd;
              .;lllodddddddddddo,
            .;llllodxxdddddddddddo.
          .;lllllodxxxxddddddddddddc.
          .;llllodxxxxxdddddddddddddd:
            .,:ldddddooooooollllllcccc,
              .',;;;;;;;;;;;;;;;;;;,.
                .,;;;;;;;;;;;;;;;,.
                 .',;;;;;;;;;;;,.
                   .,;;;;;;;;'.
                     ';;;;;'.
                      .,;


`,

	`
                      .od'
                    .:dxxxl.
                  .:ldxxxxxxc.
                .;lldxxxxxxxxd,
               ;llldxxxxxxxxxxxo'
             ,lllloxxxxxxxxxxxxxxl.
           ,llllloxxxxxxxxxxxxxxxxd:
           ,cllloxxxxxxxxxxxxxxxxxxxd,
             ,:odddddoooooooolllllllcc.
              .',;;;;;;;;;;;;;;;;;;,.
                .,;;;;;;;;;;;;;;;,.
                  ';;;;;;;;;;;;,.
                   .,;;;;;;;;'.
                     ';;;;;'.
                      .,;


`,

	`
                      .od'
                    .:dxxxl.
                  .:ldxxxxxx:
                 ;lldxxxxxxxxd'
               ,llldxxxxxxxxxxxl.
             'lllldxxxxxxxxxxxxxx:
           .clllldxxxxxxxxxxxxxxxxd,
           .cllldxxxxxxxxxxxxxxxxxxxo.
             'cdddddddoooooooolllllllc.
               ';;;;;;;;;;;;;;;;;;;,.
                .;;;;;;;;;;;;;;;;,.
                  ';;;;;;;;;;;;,.
                   .,;;;;;;;;'.
                     ';;;;;'.
                      .,;


`,

	`
                      .od'
                    .:dxxxl.
                   ;lxxxxxxx:
                 ,llxxxxxxxxxd.
               'cllxxxxxxxxxxxxc.
             .clloxxxxxxxxxxxxxxx;
           .:llloxxxxxxxxxxxxxxxxxo.
           .:lloxxxxxxxxxxxxxxxxxxxxc
             'lddddddddoooooooooollllc
               ';;;;;;;;;;;;;;;;;;;,.
                .;;;;;;;;;;;;;;;;,.
                  ';;;;;;;;;;;;'.
                   .;;;;;;;;;'
                     ';;;;;'
                      .,;


`,

	`
                      .od'
                    .:xxxxl.
                   ,lxxxxxxx;
                 'coxxxxxxxxxo.
               .:loxxxxxxxxxxxx:
             .:lloxxxxxxxxxxxxxxd'
            ;llldxxxxxxxxxxxxxxxxxl.
            :lldxxxxxxxxxxxxxxxxxxxx;
             ,ddddddddddooooooooooool;
              .,;;;;;;;;;;;;;;;;;;;'.
                .;;;;;;;;;;;;;;;;'.
                  ,;;;;;;;;;;;;'
                   .;;;;;;;;;'
                     ';;;;;'
                      .,;


`,

	`
                      .od'
                    .:xxxxc
                   ,oxxxxxxx,
                 .coxxxxxxxxxl.
               .;cdxkxxxxxxxxxx,
              ;clxkxkxkxxxxxxxxxo.
            'cclxkkkkkxkxxxxxxxxxx;
            ;coxkkkkkkkkxxkxxxxxxxxd.
             cddddddddddddooooooooooo.
              .;;;;;;;;;;;;;;;;;;;;'
                ';;;;;;;;;;;;;;;;'
                 .,;;;;;;;;;;;;.
                   .;;;;;;;;;'
                     ,;;;;;'
                      .;;


`,

	`
                      .od.
                     ;xkkkc
                   'okkkkkkx'
                 .:dkkkkkkkkkc
                ,lxkkkkkkkkkkkd'
              'clxkkkkkkkkkkkkkkc
            .:cokkkkkkkkkkkkkkkkkd'
            ,cdkkkkkkkkkkkkkkkkkkkkl
            .oddddddddddddddddooooool
              ';;;;;;;;;;;;;;;;;;;;.
               .;;;;;;;;;;;;;;;;;.
                 .;;;;;;;;;;;;;.
                   ';;;;;;;;;.
                    .,;;;;;'
                      .;;


`,

	`
                      .od.
                     ;xkkk:
                   .okkkkkkd.
                 .:dkkkkkkkkk:
                'lxkkkkkkkkkkkd.
              .:okkkkkkkkkkkkkkk;
             ;cdkkkkkkkkkkkkkkkkko.
            'lxkkkkkkkkkkkkkkkkkkkk;
            ,ddddddddddddddddddddddd;
             .;;;;;;;;;;;;;;;;;;;;;.
               .;;;;;;;;;;;;;;;;;.
                 .;;;;;;;;;;;;;.
                   ';;;;;;;;;.
                    .,;;;;;.
                      .;;


`,

	`
                      .od.
                     ;kkkk:
                   .okkkkkkd.
                  ;xkkkkkkkkk;
                .okkkkkkkkkkkko.
              .:dkkkkkkkkkkkkkkx;
             'lxkkkkkkkkkkkkkkkkkl.
            .okkkkkkkkkkkkkkkkkkkkx'
            lddddddddddddddddddddddd.
             .;;;;;;;;;;;;;;;;;;;;,.
               ';;;;;;;;;;;;;;;;,.
                 ';;;;;;;;;;;;;.
                  .,;;;;;;;;;.
                    .;;;;;;.
                      .;;


`,

	`
                      .dd.
                     ;kkkk:
                   .dkkkkkkd.
                  ;kkkkkkkkkx,
                .okkkkkkkkkkkkl.
               ;xkkkkkkkkkkkkkkd'
             .okkkkkkkkkkkkkkkkkk:.
            .xkkkkkkkkkkkkkkkkkkkko.
            dddddddddddddddddddddddl
             ':::::::::::::::::::;'
               ,::::::::::::::::,
                .,::::::::::::,.
                  .,::::::::;.
                    .;::::;.
                      .;;


`,

	`
                      .dd.
                     ;kkkk:
                   .dkkkkkko.
                  ,kkkkkkkkkx,
                .okkkkkkkkkkkk:.
               ,kOkkkkkkkkkkkkko'.
             .oOOOOOOOOOOkkkkkkkd,.
            'kOOOOOOOOOOOOOOkkkkkk:.
           'ddddddddddddddxxxxxxxxx;
            .,:::::::::::::::::::;.
              .;::::::::::::::::'
                .;::::::::::::'
                  .;::::::::,.
                    .;::::;.
                      .;;


`,

	`
                      .dx.
                     :kkkk:
                   .dkkkkkko.
                  :kOOOkkkkkd'.
                .dOOOOOOOOOOOx;.
               :kOOOOOOOOOOOOOkc'.
             .dOOOOOOOOOOOOOOOOOl''.
            :OOOOOOOOOOOOOOOOOOOOd,.
           cdddddddddddxxxxxxxxxxxd.
            .;:::::::::::::::::::,.
              .;:::::::::::::::;.
                .;::::::::::::'
                  .;::::::::,
                    .;::::;.
                      .;;


`,

	`
                      .dd.
                     :kOOk;
                   .xOOOOOkl.
                  cOOOOOOOOOo'.
                'xOOOOOOOOOOOd,..
               lOOOOOOOOOOOOOOx;'.
             ,kOOOOOOOOOOOOOOOOk:''.
           .dOOOOOOOOOOOOOOOOOOOOc'.
          .odddddddddddxxxxxxxxxxxc
            .;:::::::::::::::::::'
              .;:::::::::::::::;.
                .;:::::::::::;.
                  .;::::::::'
                    .;::::;.
                      .;;


`,

	`
                      .dd.
                     ckOOk;
                   ,kOOOOOkc.
                 .oOOOOOOOOOc'.
                ;kOOOOOOOOOOOl''.
              .dOOOOOOOOOOOOOOd,''.
             cOOOOOOOOOOOOOOOOOd,''.
           ,kOOOOOOOOOOOOOOOOOOOx;''
          'dddddddddddxxxxxxxxxxxx'
            ':::::::::::::::::::;.
              '::::::::::::::::'
                '::::::::::::;.
                  .:::::::::.
                    ':::::;.
                      .;;


`,

	`
                      .dd.
                    .lOOOx,.
                   ;kOOOOOk:.
                 .dOOOOOOOOk:,.
                ckOOOOOOOOOOkc,'.
              ,kOOOOOOOOOOOOOOc,,'.
            .oOOOOOOOOOOOOOOOOOl,,,'.
           cOOOOOOOOOOOOOOOOOOOOo,,'.
          :oodddddddddxxxxxxxxxkkl.
           .,:::::::::::::::::::;.
             .,::::::::::::::::.
                '::::::::::::,
                  ':::::::::.
                    ':::::,
                      .;;


`,

	`
                      .xd.
                    .lOOOx,.
                   :kOOOOOx;'.
                 .xOOOOOOOOx;,.
               .lOOOOOOOOOOOx;,,.
              :kOOOOOOOOOOOOOk;,,,.
            'xOOOOOOOOOOOOOOOOk:,,,,.
          .oOOOOOOOOOOOOOOOOOOOk:,,,.
          loooddddddddxxxxxxxxkkx;.
           .,:::::::::::::::::::,.
             .,:::::::::::::::;.
               .,::::::::::::,
                  ':::::::::.
                    ':::::,
                      .:;


`,

	`
                      'xd.
                    .oOOOx,.
                   ckOOOOOx;'.
                 ,xOOOOOOOOd,,'
               .dOOOOOOOOOOOd,,,.
             .lOOOOOOOOOOOOOOd,,,,.
            :kOOOOOOOOOOOOOOOOd,,,,,.
          'xOOOOOOOOOOOOOOOOOOOd,,,,.
         .loooodddddddxxxxxxxkkko'.
           .;:::::::::::::::::::'.
             .;:::::::::::::::;.
               .,::::::::::::'
                  '::::::::;.
                    ':::::,
                      ':;


`,

	`
                      'xd.
                    .okkkd,.
                   ckkkkkkd,,.
                 ;kkOOOOOOOo,,'.
               'xOOOOOOOOOOOl,,,'.
             .dOOOOOOOOOOOOOOc,,,,'
           .lOOOOOOOOOOOOOOOOkc,,,,,'
          :kOOOOOOOOOOOOOOOOOOk:,,,,'.
         'loooooddddddxxxxxxkkkx;'..
           .;::::::::::::::::::;..
             .;:::::::::::::::,.
               .,:::::::::::;'.
                 .,::::::::;.
                    ':::::,
                      ':;


`,

	`
                      'xd.
                    .okkkd,.
                  .lkkkkkko;,.
                 :kkkkkkkkkc;;,.
               ;xkkkkkkkkkkk:;;;,.
             'xkkkkkkkkOOOOOk:;;;;,.
           .okkkkkkOOOOOOOOOOx;;;;;;'.
         .lkkkOOOOOOOOOOOOOOOOd;;;;,'.
         ;lloooooddddxxxxxxkkkko,'..
           .;::::::::::::::::::,..
             .;::::::::::::::;'..
               .,:::::::::::;..
                 .,::::::::,.
                   .,:::::'
                      ':;


`,

	`
                      'xo.
                    .okkko,.
                  .okkkkkkc;,.
                 ckkkkkkkkk:;;,.
               :kkkkkkkkkkkx;;;;,.
             ;xkkkkkkkkkkkkkd;;;;;,.
           ,dkkkkkkkkkkkkkkkkl;;;;;;,.
         .dkkkkkkkkkkkkkkkkkkkc;;;;;,.
         ;llooooodddddxxxxxkkkk;,'..
           .::::::::::::::::::;...
             .;::::::::::::::,...
               .;:::::::::::,..
                 .,::::::::'.
                   .,:::::'
                      ':;


`,

	`
                      'xo.
                    'dkkkl;.
                  .okkkkkk:;;.
                .lkkkkkkkkx;;;;.
              .ckkkkkkkkkkko;;;;;.
             :kkkkkkkkkkkkkkc;;;;;;.
           ;xkkkkkkkkkkkkkkkk:;;;;;;;.
         ,xkkkkkkkkkkkkkkkkkkx;;;;;;,'.
        .;llloooooddddxxxxkkkkl;,'...
           '::::::::::::::::::,....
             .;::::::::::::::'...
               .;:::::::::::'..
                 .,::::::::'.
                   .,:::::.
                      ':;


`,

	`
                      'do.
                    'dkkkl;.
                  .okkkkkx:;;.
                .lkkkkkkkkd;;;;.
              .ckkkkkkkkkkkc;;;;;.
            .ckkkkkkkkkkkkkx:;;;;;;.
          .:kkkkkkkkkkkkkkkkd;;;;;;;;.
         :kkkkkkkkkkkkkkkkkkkc;;;;;;,,.
        .;llllooooddddxxxxxkkx:;,'...
           ':::::::::::::::::;.....
             .;:::::::::::::;....
               .;::::::::::;...
                 .,:::::::;..
                   .,::::;..
                      ':;


`,

	`
                      ,dl.
                    'dkkkc;.
                  'dkkkkkx::;.
                .okkkkkkkkl:::;.
              .lkkkkkkkkkkx:::::;.
            .lkkkkkkkkkkkkko::::::;'
          .ckkkkkkkkkkkkkkkkc:::::::;'
        .ckkkkkkkkkkkkkkkkkkx::::::;;,.
        .;cllllooooddddxxxkkkl;;,'...
           ':::::::::::::::::,.....
             .;:::::::::::::'....
               .;::::::::::,...
                 .,:::::::,..
                   .,::::;..
                      ':;


`,

	`
                      ,dl.
                    'dkkxc;.
                  'dkkkkkd:::'
                .okkkkkkkkc::::'
              .okkkkkkkkkkd::::::'
            .okkkkkkkkkkkkkc:::::::'
          .lkkkkkkkkkkkkkkkd:::::::::'.
        .okkkkkkkkkkkkkkkkkkl:::::::;,'
        .;cclllloooddddxxxkkx:;,,'...
          .'::::::::::::::::;......
             .;::::::::::::;.....
               .;::::::::::'...
                 .,:::::::'...
                   .,::::,..
                      ':;


`,

	`
                      ,dl.
                    ,dxkx::'
                  ,dxkkkko:::'
                'dxkkkkkkx:::::'
              'oxkkkkkkkkkl::::::'
            'dxxkkkkkkkkkkx::::::::,
          'oxxkkkkkkkkkkkkkl:::::::::,.
        'dxxkkkkkkkkkkkkkkkd::::::::;;,
        .;cccllloooodddxxxkkc:;,,'....
          .'::::::::::::::::'.......
             .;::::::::::::,......
               .;:::::::::;.....
                 .,:::::::....
                   .,::::,..
                      ':,


`,

	`
                      ,dl.
                    ,dxxx::'
                  ,dxxxxxl:::'
                'dxxxxxxxd:::::'
              ,dxxxxxxxxxxc::::::,.
            ,dxxxxxxxxxxxxo::::::::,.
          ,dxxxxxxxxxxxxxxx::::::::::,.
        ,dxxxxxxxxxxxxxxxxxl:::::::::;,.
        .,cccllllooodddxxxkd::;,,''...
           ';::::::::::::::;........
             .;::::::::::::.......
               .;:::::::::,.....
                 .,::::::;....
                   .,::::'..
                      ';,


`,

	`
                      ,dl'
                    ,dxxdc:'
                  ,dxxxxxcccc,
                ,dxxxxxxxoccccc,
              ,dxxxxxxxxxdccccccc,.
            ,dxxxxxxxxxxxxccccccccc;.
          ;dxxxxxxxxxxxxxxocccccccccc;.
        ,dxxxxxxxxxxxxxxxxdccccccccc:;;.
        .,:ccclllooodddxxxxc::;;,''...
           ';::::::::::::::'........
             .;:::::::::::,.......
               .;:::::::::......
                 .,::::::,....
                   .,::::...
                      ';,


`,

	`
                      ,dc'
                    ,dxxocc'
                  ;dxxxxxcccc,
                ,dxxxxxxxlccccc,
              ;dxxxxxxxxxoccccccc,.
            ;dxxxxxxxxxxxdccccccccc;.
          ;dxxxxxxxxxxxxxxccccccccccc;.
        ;dxxxxxxxxxxxxxxxxlccccccccc::;'
        .,::cclllooodddxxxoc::;;,''...
           ';:::::::::::::;.........
             .;:::::::::::........
               .,::::::::,......
                 .,::::::'....
                   .':::;...
                      ';'


`,

	`
                      ,dc'
                    ,dxxocc'
                  ;dxxxxdcccc,
                ;dxxxxxxxcccccc,
              ;dxxxxxxxxxlccccccc,.
            ;xxxxxxxxxxxxlccccccccc;.
         .:xxxxxxxxxxxxxxoccccccccccc;.
        :dxxxxxxxxxxxxxxxdccccccccccc::'
         '::cccllloodddxxxcc::;;,''...
           .;;;;;;;;;;;:::..........
             .;;;;;;;;;;;,........
               .,;;;;;;;;'......
                 .,;;;;;;.....
                   .';;;,...
                      .;'


`,

	`
                      ,dc'
                    ,dxxlcc,
                  ;dxxxxocccc,
                ;dxxxxxxocccccc,
              ;dxxxxxxxxdcccccccc,.
           .:xxxxxxxxxxxxcccccccccc;.
         .:xxxxxxxxxxxxxxlccccccccccc;.
        :dxxxxxxxxxxxxxxxlccccccccccc::,
         ':::ccllloodddxxocc::;;,,''...
           .;;;;;;;;;;:::,..........
             .;;;;;;;;;;;.........
               .,;;;;;;;;.......
                 .,;;;;;,.....
                   .';;;,...
                      .;'


`,

	`
                      ,dc'
                    ,dxxllc,
                  ;dxxxxllllc,
                ;dxxxxxxolllllc,
             .;dxxxxxxxxolllllllc;.
           .:dxxxxxxxxxxollllllllll;.
         .:dxxxxxxxxxxxxollllllllllll;.
        :ddxxxxxxxxxxxxxdlllllllllllcc:,
         .;::cccllooddxxdlcc::;;,,''...
           .;;;;;;;;;;::;'...........
             .,;;;;;;;;;,.........
               .,;;;;;;;,.......
                 .';;;;;'.....
                    ';;;'...
                      .;.


`,

	`                       ..
                     .:Ok;.
                   .:dxOkll;.
                 .:xxxxOxllll;.
               .:xxxxxxkxllllll;.
             .:dxxxxxxxkdllllllll;.
           .cxxxxxxxxxxkdllllllllll:.
         .cxxxxxxxxxxxxxollllllllllll:.
        cdxxxxxxxxxxxxxxllllllllllllllc;
         .;:clloddxkkkkxlloollc::;;,''.
           .;;;;;;;;:clo;;,'.........
             .;;;;;;;;;;'..........
               .,;;;;;;:'........
                 .,;;;;:,......
                   .,;;:,...
                     .':;


`,

	`                       ..
                     .lXX:.
                   .cxxKKll:.
                 .lxxxx00llll:.
               .lxxxxxx0Ollllll:.
             .cxxxxxxxxOkllllllll:.
           .lxxxxxxxxxxOxllllllllll:.
         .lxxxxxxxxxxxxkdllllllllllll:.
        cxxxxxxxxxxxxxxxollllllllllllll:
         ':clodxxkO00OOOoooddoolc::;,'.
           .;;;;;;;:codxc:;,'........
             .;;;;;;;;;:'..........
               .;;;;;;;:,........
                 .;;;;;:;......
                   .;;;c:....
                     .,cc..
                       .

`,

	`                       ..
                     'oWMl'
                   'oxxNMllc'
                 'oxxxxXNllllc'
               .oxxxxxxKXllllllc.
             .lxxxxxxxx00llllllllc.
           .oxxxxxxxxxxOOllllllllllc.
         .lxxxxxxxxxxxxOxllllllllllll:.
        lxkxxxxxxxxxxxxkdllllllllllllll:.
         'clodxkO0KKK000ddddxxdolcc:;,.
           ';;;;;;:lodxOlc:;,'.......
             ';;;;;;;;;:'..........
               .;;;;;;;:;........
                 .;;;;;c:......
                   .;;;cc....
                     .;ll..
                       .

`,

	`                      ....
                     ;kMMk,
                   ;dxxMMdll,
                 ,dxxxxWMollll,
               'dxxxxxxNWllllllc'
             'oxxxxxxxxXXllllllllc'
           'oxxxxxxxxxx00llllllllllc.
         'oxxxxxxxxxxxxOkllllllllllllc.
        okkxxxxxxxxxxxxkdllllllllllllllc.
        .,lodkO0KXXXXKKKxxxxxxxxdolc:,'
          .';;;;;:lodkOKdolc:,'......
             ';;;;;;;;;:'..........
               ';;;;;;;:;........
                 ';;;;;cc......
                   ';;;ll'...
                     ';ll,.
                       .

`,

	`                      ....
                    .cKMMK:.
                  .:dxOMMOll;.
                 :dxxxxMMxllll;.
               ;dxxxxxxWMollllll,
             ,dxxxxxxxxNWllllllllc'
           ,dxxxxxxxxxxKKllllllllllc'
         'oxxxxxxxxxxxxOOllllllllllllc.
       .okkxxxxxxxxxxxxkdllllllllllllloc.
        .;ldkO0XNNNNNXXXxxxkkkkkxdol:;'.
          .,;;;;:codkOKXxdolc:,'.....
            .,;;;;;;;;;:'..........
              .,;;;;;;;c:........
                .,;;;;;cl......
                  .,;;;ll,....
                    .,:ll;..
                      ...

`,

	`                      ....
                    .oNMMWc.
                  .lxxKMMXll:.
                .cxxxxOMMOllll:.
               :xxxxxxxMMxllllll;.
             ;dxxxxxxxxWMlllllllll,
           ,dxxxxxxxxxxXXlllllllllll,
         ,oxxxxxxxxxxxx0Ollllllllllllc'
       .dkkkxxxxxxxxxxxkdllllllllllllooc.
        .:oxO0XNWWWWWNNNkkkkOOOOkxdlc:,.
          .,;;;:codk0KNWOkxolc;,'.....
            .,;;;;;;;;;c,...........
              .,;;;;;;;c:.........
                .,;;;;;ll'......
                  .;;;:ll;....
                    .;cllc..
                      ...

`,

	`                      .....
                    ,xMMMMx,
                  'oxxNMMWllc'
                .oxxxxKMMXllllc.
              .cxxxxxxOMMOllllll:.
            .:dxxxxxxxxWMdllllllll;.
           ;dxxxxxxxxxxXNlllllllllll,
         ,dxxxxxxxxxxxx0Ollllllllllllc'
       .dkkkkkkkxxxxxxxkdllllllllllooool.
        .cdkOKNMMMMMWMWWOOOOOOOOOkxol:;.
          .,;;;cldkOKNWMOOkxolc;,'....
            .;;;;;;;;;:o:,..........
              .;;;;;;;;c:.........
                .;;;;;;ll'......
                  .;;;:ll:....
                    .;clll'.
                      ...

`,

	`                     ......
                    :0MMMM0:.
                  ;dxxMMMMdll,
                ,dxxxxXMMWllllc'
              .oxxxxxx0MM0llllllc.
            .cxxxxxxxxxMMxllllllll:.
           :xxxxxxxxxxxNWlllllllllll;
         ,dxxxxxxxxxxxx00llllllllllllc'
       .xkkkkkkkkkkkkkkOdooooooooooooool.
        .ldk0XWMMMMMMMMMOOOOOOOOOOkdoc;.
          .;;;:ldxOKNWMMOOOkxol:;'....
            .;;;;;;;;:ldc;,.........
              .;;;;;;;;c:.........
                ';;;;;;ll,......
                  ';;;cll:....
                   .':llll,.
                     .....

`,

	`                     ......
                   .lXMMMMNc.
                 .cxxOMMMMOll:.
                ;dxxxxWMMMollll,
              'oxxxxxxKMMXllllllc.
            .lxxxxxxxxkMMkllllllll:.
           cxxxxxxxxxxxNWlllllllllll;
         ;dxxxxxxxxxxxx00lllllllllllll,
       .xkkkkkOOOOOOOOO0xdddoooooooooool.
        'oxOKWMMMMMMMMMMOOOOOOOOOOOxol:.
          .;;;coxOKNWMMMOOOOkdoc:,'...
            ';;;;;;;:ldkoc;'........
              ';;;;;;;;::.........
                ,;;;;;:cl,......
                 .,;;;cllc.....
                   .,:llll;..
                     .....

`,

	`                     .......
                   'dWMMMMMo'
                 .lxxKMMMMXllc.
                cxxxxxMMMMxllll:.
              ,dxxxxxxXMMNlllllll'
            .oxxxxxxxxOMMOllllllllc.
          .cxxxxxxxxxxxNWollllllllll:.
         ;dxxxxxxxxxxxx00lllllllllllll,
       .xkkkOOOOO00000KKxxdddddddooooool.
        ,dk0XWMMMMMMMMMM0000000000Okdl:'
          .;;:ldk0XWMMMM0000Okdlc;'...
            ';;;;;;:ldk0dlc;'.......
             .,;;;;;;;;:;.........
               .,;;;;;:cc,.......
                 .;;;;cccc'....
                   .;:cccc:..
                     .....

`,

	`                    ........
                   :kMMMMMMk;
                 'dxxNMMMMWllc'
               .lxxxxOMMMMOllllc.
              ;dxxxxxxNMMWollllll;
            'oxxxxxxxxOMM0llllllllc.
          .lxxxxxxxxxxxWWollllllllll:.
         :xxxxxxxxxxxxx0Olllllllllllll,
       .xkkkOOO000KKKKXXkkxxxxdddddooool.
        ;dOKNMMMMMMMMMMM00000000000Oxoc,
          ';;coxOKWMMMMM00000Oxol:,'..
           .,;;;;;;cdk0Xkdl:,'......
             .,;;;;;;;;:;..........
               .;;;;;;:cc,.......
                 .;;;;cccc'....
                   ';cccccc'.
                     .....

`,

	`                    ........
                  .lKMMMMMMXc.
                 ;dxxWMMMMMdll,
               .oxxxx0MMMMKllllc.
              cxxxxxxxWMMMollllll:
            ,dxxxxxxxx0MM0llllllllc'
          .oxxxxxxxxxxxWWollllllllllc.
         :xxxxxxxxxxxxx0Olllllllllllll;
       .xkkOOO00KKKXXXNNOkkkkxxxddddoool.
        :xOXWMMMMMMMMMMM000000000000kdl;.
          ';:ldkKNMMMMMM000000kxoc;'..
           .,;;;;;cox0XW0kdl:,.......
             .;;;;;;;;;c;..........
               .;;;;;;:cc,.......
                 ';;;:cccc,....
                  .,;cccccc'..
                    .......

`,

	`                    .........
                  .dNMMMMMMWl'
                .cxxkMMMMMMkll:.
               ,dxxxxXMMMMNlllll,
             .lxxxxxxxWMMMxllllll:.
            ;dxxxxxxxx0MMKlllllllll,
          .oxxxxxxxxxxxWWollllllllllc.
         :xxxxxxxxxxxxxOklllllllllllll;
       .xkkOO00KKXXNNWWW0OOOkkkxxxdddool.
       .ck0NMMMMMMMMMMMM000000000000Odl;.
         .,;coxOXWMMMMMM0000000kdl:,..
           .;;;;;:lxOXWM00koc;'......
             .;;;;;;;;:l:'.........
               ';;;;;;:cc,.......
                .,;;;:cccc,.....
                  .;:cccccc,..
                    .......

`,

	`                   ..........
                  ;xWMMMMMMMx,
                .oxx0MMMMMM0llc.
               :xxxxxNMMMMWollll;
             .oxxxxxxkMMMMkllllllc.
            :xxxxxxxxxKMMXlllllllll;
          .dxxxxxxxxxxxNWollllllllllc.
         cxxxxxxxxxxxxk0kollllllllllll;
       .xkkO00KKXXNNWMMM0000OOkkxxxddool.
       .lkKWMMMMMMMMMMMM000000000000Oxo:.
         .,;ldkKNMMMMMMM0000000Oxoc;'..
           .;;;;;cdkKNMM00Oxoc,......
             ';;;;;;;;cdl;'........
              .,;;;;;;:::,.......
                .;;;;:::::,.....
                  .;:::::::;..
                    .......

`,

	`                   ..........
                  cOMMMMMMMM0:.
                'dxxXMMMMMMNlll'
               lxxxxxWMMMMWdllll:.
             'dxxxxxxOMMMMOlllllll'
            cxxxxxxxxxKMMXlllllllll:
          'dxxxxxxxxxxxNWolllllllllll'
         cxxxxxxxxxxxkOKkooollllllllll:.
       .xkOO00KXXNWWMMMM00000OOkkxxddool.
       .oOXWMMMMMMMMMMMM0000000000000kdc.
         .;:lxOXWMMMMMMM00000000kdl;'..
           .;;;;:lx0NMMM000Odl:'.....
             ,;;;;;;;coko:,........
              .;;;;;;;:::,........
                ';;;;:::::;.....
                  ,;:::::::;'.
                   .........

`,

	`                   ..........
                 .oXMMMMMMMMXl.
                ;dxxNMMMMMMWoll;
              .oxxxxkWMMMMMxllllc.
             ,dxxxxxxOMMMM0lllllll,
           .lxxxxxxxxxKMMXlllllllll:.
          ,dxxxxxxxxxxxNNolllllllllll'
        .lxxxxxxxxxkkO0Kkddoolllllllll:.
       .xkOO0KXXNWMMMMMM0000000Okkxxddol.
       .d0XMMMMMMMMMMMMM0000000000000Odc.
         .;cok0NMMMMMMMM00000000Oxo:,..
           ';;;;cdOKWMMM0000koc,.....
            .,;;;;;;:lx0dl:'........
              .;;;;;;;;::,........
                ,;;;;:::::,.....
                 .;;:::::::;'..
                   .........

`,

	`                  ............
                 ,xNMMMMMMMMWo,
                cxxkWMMMMMMWxll:.
              'dxxxxOMMMMMMklllll'
             :xxxxxxx0MMMM0lllllll;
           .oxxxxxxxxxKMMXlllllllllc.
          ;xxxxxxxxxxxxXNolllllllllll,
        .lxxxxxxxxkkO0KXOxxdoollllllll:.
       .xkO00KXNWMMMMMMMK0K0KK00Okkxddol.
       .d0NMMMMMMMMMMMMMK0K00K00K00KK0xl.
         .;cdOKWMMMMMMMMK0K00K000kdc;'.
           ,;;;:lx0NMMMMK0K0Oxl:'....
            .;;;;;;;cdOKkoc;'.......
              ';;;;;;;;:;'........
               .;;;;;;::::,'.....
                 .;;:::::::;'..
                   .........

`,

	`                  ............
                 :kWMMMMMMMMMk:
               .oxxOMMMMMMMMOllc.
              ,dxxxxOMMMMMM0lllll,
             cxxxxxxx0MMMMKlllllll:
           .dxxxxxxxxxKMMXlllllllllc.
          ;xxxxxxxxxxxxXXllllllllllll,
        .lxxxxxxxkkO0KXNOkkxdoolllllll:.
       .xkO0KXXNWMMMMMMMKKKKKKK00Okxxdol.
       'xKWMMMMMMMMMMMMMKKKKKKKKKKKKK0ko'
         .:lx0XMMMMMMMMMKKKKKKKKKOxl:'.
          .,;;;cdOXWMMMMKKKKKkdc,....
            .;;;;;;:ok0NOxo:,.......
              ,;;;;;;;;:;'........
               .;;;;;;:::;,'.....
                 ';;;::::::;,..
                  ..........

`,

	`                  ............
                .o0MMMMMMMMMM0c.
               'dxx0MMMMMMMMKlll'
              :xxxxx0MMMMMMKlllll;
            .oxxxxxxxKMMMMKlllllllc.
           'dxxxxxxxxxKMMKllllllllll'
          :xxxxxxxxxxxxKKllllllllllll;
        .oxxxxxxxkO0KXNW0Okkxdoollllllc.
       .xkO0KXNWMMMMMMMMKKKKKKKK0OOkxdol.
       'kXWMMMMMMMMMMMMMKKKKKKKKKKKKKKOo,
         ':okKWMMMMMMMMMKKKKKKKKK0ko:,.
          .;;;:lx0NMMMMMKKKKK0xl;'....
            ';;;;;;lxOXMKOdl;'......
             .,;;;;;;;;c:'.........
               ';;;;;;;;;;,'.....
                .,;;;;;;;;;;,'..
                  ...........

`,

	`                 ..............
                .dXMMMMMMMMMMNo'
               ;xxxKMMMMMMMMXoll;
              lxxxxxKMMMMMMKlllllc.
            .dxxxxxxxKMMMMKlllllllc.
           ,dxxxxxxxxx0MMKllllllllll,
          cxxxxxxxxxxxxK0llllllllllll:
        .oxxxxxxkO0KXNWMK0OOkxdollllllc.
       .xkO0KXNWMMMMMMMMKKKKKKKKK0Okxdol.
       ,OXMMMMMMMMMMMMMMKKKKKKKKKKKKKKOd,
         ,cdOXWMMMMMMMMMKKKKKKKKKKOdc,.
          .;;;:okXWMMMMMKKKKKKOdc,....
            ';;;;;:okKWMK0koc,......
             .;;;;;;;;:oc'.........
               ,;;;;;;;;;;,'.....
                .;;;;;;;;;;;,'..
                  ...........

`,

	`                 ..............
                ;kNMMMMMMMMMMWd;
               lxxxXMMMMMMMMNdll:
             .oxxxxxXMMMMMMXollllc.
            .dxxxxxxxKMMMMKllllllll'
           ;xxxxxxxxxx0MM0llllllllll,
          lxxxxxxxxxxxxK0olllllllllll:
        .oxxxxxxO0KXNWMMKK0Okkxdolllllc.
       .xkO0XNWMMMMMMMMMKKKKKKKKK0Okxdol.
       ;ONMMMMMMMMMMMMMMKKKKKKKKKKKKKK0x;
         ,lx0NMMMMMMMMMMKKKKKKKKKK0xl;.
          .;;;cx0NMMMMMMKKKKKK0xl;....
           .,;;;;;lx0NMMKK0xl;'.....
             .;;;;;;;;cxo;'........
              .;;;;;;;;;;;'.......
                ';;;;;;;;;;;,'..
                  ...........

`,

	`                 ..............
                lOWMMMMMMMMMMMkc.
              .oxxkNMMMMMMMMWxllc.
             .dxxxxxXMMMMMMNolllll'
            ,xxxxxxxxKMMMMKllllllll,
           :xxxxxxxxxxOWMOllllllllll;
          lxxxxxxxxxxxOKOdollllllllllc
        .oxxxxxkO0KNWMMMKKK0Okxdolllllc.
       .xkOKXNWMMMMMMMMMKKKKKKKKKK0Okdol.
       :0WMMMMMMMMMMMMMMKKKKKKKKKKKKKKKx:
        .;lxKWMMMMMMMMMMKKKKKKKKKKKko:'
          ';;:lkKWMMMMMMKKKKKKKko:'...
           .;;;;;:okKWMMKKKkdc,......
             ';;;;;;;:oOdc,........
              .;;;;;;;;;;,'.......
                ,;;;;;;;;;;,,'..
                 ............

`,

	`                 ...............
               .d0MMMMMMMMMMMMKl.
              'dxxOWMMMMMMMMWklll'
             ;xxxxxxXMMMMMMNdlllll;
            :xxxxxxxxKMMMMKllllllll;
           cxxxxxxxxxxOWWkllllllllll:
          oxxxxxxxxxxk0XOxdllllllllllc.
        .dxxxxxk0KXNMMMMKKKK0Okxdollllc.
       .xkOKXNMMMMMMMMMMKKKKKKKKKK0Okxol.
       :0WMMMMMMMMMMMMMMKKKKKKKKKKKKKKKk:
        .:okKWMMMMMMMMMMKKKKKKKKKKKOdc'.
          ,;;:dONMMMMMMMKKKKKKK0xc,...
           .;;;;;cx0NMMMKKK0xl;'.....
             ,;;;;;;;lxKko;'.......
              ';;;;;;;;;,,'.......
               .;;;;;,,,,,,,,'...
                 ............

`,

	`                ................
               ,xXMMMMMMMMMMMMXo,
              ;xxx0WMMMMMMMMWOlll;
             cxxxxxkXMMMMMMNdlllll:
            cxxxxxxxx0MMMM0llllllll:
           lxxxxxxxxxxkNNxllllllllllc.
         .oxxxxxxxxxk0KN0kxolllllllllc.
        .dxxxxkO0XNWMMMMXXXXK0Oxdollllc.
       .dk0KXWMMMMMMMMMMXXXXXXXXXXK0kxol.
       cKMMMMMMMMMMMMMMMXXXXXXXXXXXXXXKOc
        .:dOXMMMMMMMMMMMXXXXXXXXXXX0xc,.
          ,;;cx0WMMMMMMMXXXXXXXKkl;...
           .;;;;:okKWMMMXXXKOd:'.....
            .;;;;;;;:dON0xc,........
              ,;;;;;;;;;,''.......
               .;;;;,,,,,,,,''...


`,

	`                ................
               ckNMMMMMMMMMMMMNx:
              lxxx0WMMMMMMMMM0lllc
             oxxxxxkXMMMMMMNdlllllc.
            oxxxxxxxx0WMMM0llllllllc.
          .oxxxxxxxxxxkXXdllllllllllc.
         .dxxxxxxxxkOKXWKOkdolllllllll.
        .dxxxxk0KXWMMMMMXXXXXKOkxolllll.
       .dk0KNWMMMMMMMMMMXXXXXXXXXXK0kxdl.
       lKMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXOl
        .cx0NMMMMMMMMMMMXXXXXXXXXXXKxl;.
         .;;;lkXMMMMMMMMXXXXXXXKOd:'..
           ';;;;cdONMMMMXXXX0xl,.....
            .;;;;;;;lkKWKOo:'.......
             .;;;;;;;;;:;''........
               ';;;,,,,,,,,,''...


`,

	`                .................
              .oOWMMMMMMMMMMMMWOl.
             .oxxxKWMMMMMMMMMKollc.
            .dxxxxxkXMMMMMMNxllllll.
           .dxxxxxxxx0WMMWOlllllllll.
          .dxxxxxxxxxxxXKolllllllllll.
         .dxxxxxxxxk0XNMXKOxdlllllllll.
        .dxxxxO0KNWMMMMMXXXXXK0Oxdollll.
       .dk0KNWMMMMMMMMMMXXXXXXXXXXX0Oxdl.
       oXMMMMMMMMMMMMMMMXXXXXXXXXXXXXXX0l
        .lxKWMMMMMMMMMMMXXXXXXXXXXXKko:.
         .;;:oONMMMMMMMMXXXXXXXX0xc,...
           ,;;;;lxKWMMMMXXXXKko:'....
            ';;;;;;:dONMXKxc,.......
             .;;;;;;;;;lc'.........
              .,;;;,,,,,,,,'''...


`,

	`               ..................
              .x0WMMMMMMMMMMMMW0o'
             .dxxxKWMMMMMMMMMKdlll'
            .dxxxxxkXMMMMMMXxllllll'
           .dxxxxxxxxONMMWklllllllll'
          .dxxxxxxxxxxxXKolllllllllll.
         .dxxxxxxxkOKNWMXX0kxollllllll.
        .dxxxxOKXWMMMMMMXXXXXXKOkdollll.
       .dk0XNMMMMMMMMMMMXXXXXXXXXXXKOxdl.
       oNMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXKo
        'okXMMMMMMMMMMMMXXXXXXXXXXXXOd:.
         .;;cx0WMMMMMMMMXXXXXXXXKkl;...
          .;;;;:oOXMMMMMXXXXX0xc,....
            ,;;;;;;lxKWMXXOo;.......
             ';;;;;;;;:dl;.........
              .;;;;,,,''''''''...


`,

	`               ..................
              :xKWWWWWWWWWWWWWWKd:
             ;xxxkKWWWWWWWWWWKdlll;
            ;xxxxxxkXWWWWWWXxllllll;
           ,xxxxxxxxxOXWWNklllllllll,
          'xxxxxxxxxxxkXKdlllllllllll,
         'xxxxxxxxO0XWMMXXK0kdllllllll'
        .dxxxkOKNWMMMMMMXXXXXXX0Oxollll'
       .xk0XWMMMMMMMMMMMXXXXXXXXXXXKOkdl.
       dNMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXKd
        ,oOXMMMMMMMMMMMMXXXXXXXXXXXX0xc'
         .;;lkXMMMMMMMMMXXXXXXXXXOd:...
          .;;;;cx0WMMMMMXXXXXKkl;.....
            ;;;;;;:oONMMXXKxc'......
             ,;;;;;;;;lkdc'........
              ';;;;,,,''''''''....


`,

	`               ..................
              okXWWWWWWWWWWWWWWXxc
             lxxxkKWWWWWWWWWWKxlllc
            cxxxxxxkKWWWWWWKdllllllc
           :xxxxxxxxxkXWWXxlllllllll:
          ;xxxxxxxxxxxOXKxollllllllll;
         ,xxxxxxxxOKNMMMXXXKOxolllllll,
        'xxxxk0XNMMMMMMMXXXXXXXKOxdllll,
       .xk0XWMMMMMMMMMMMXXXXXXXXXXXX0kdl'
       xNMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXKd
        ;d0NMMMMMMMMMMMMXXXXXXXXXXXXKxl,
         ';;oONMMMMMMMMMXXXXXXXXXKxc'..
          .;;;;lkKMMMMMMXXXXXXOd:'....
           .;;;;;;cxKWMMXXXOo;.......
            .;;;;;;;;cd0kl,.........
              ,;;;;,,,,'''''''....


`,

	`              ....................
             .dOXWWWWWWWWWWWWWWXkl.
            .dxxxkXWWWWWWWWWWKxllll.
            oxxxxxxkKWWWWWWKdllllllc.
           lxxxxxxxxxxKWW0dlllllllllc
          cxxxxxxxxxxk0N0kdllllllllll:
         ;xxxxxxxk0XWMMMXXXX0kdlllllll;
        ,xxxxO0XWMMMMMMMXXXXXXXX0kdllll,
       .xk0XWMMMMMMMMMMMXXXXXXXXXXXX0kdl'
       kWMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXXx
        :x0WMMMMMMMMMMMMXXXXXXXXXXXXXko,
         ,;:d0WMMMMMMMMMXXXXXXXXXXkl,..
          ';;;:oONMMMMMMXXXXXXKxc'....
           .;;;;;;lkNMMMXXXKx:'......
            .;;;;;;;;lkN0x:'........
             .;;;;;,,,,,''''''.....


`,

	`              ....................
             ;x0XWWWWWWWWWWWWWWNOo;
            'xxxxOKWWWWWWWWWWKxllll,
           .dxxxxxxk0NWWWWN0dlllllll.
           oxxxxxxxxxx0NNOollllllllll.
          lxxxxxxxxxxOKNKOxolllllllllc
         cxxxxxxxOKNMMMMXXXXKOxollllll:
        ;xxxxOKNWMMMMMMMXXXXXXXX0kxolll;
       'xOKNWMMMMMMMMMMMXXXXXXXXXXXX0kdl,
       kWMMMMMMMMMMMMMMMXXXXXXXXXXXXXXXXx
        cxKWMMMMMMMMMMMMXXXXXXXXXXXXXOo;
         ,;cxKWMMMMMMMMMXXXXXXXXXXOo;..
          ,;;;cd0WMMMMMMXXXXXXXko;....
           ';;;;;:d0WMMMXXXXkl,......
            ';;;;;;;:d0WXOl,........
             .;;;;;;,,,;,'''''.....


`,

	`              ....................
             lk0NWWWWWWWWWWWWWWNOdc
            :xxxxOKNWWWWWWWWWKkllll:
           'xxxxxxxx0NWWWWNOolllllll,
          .dxxxxxxxxxxOXKklllllllllll.
          oxxxxxxxxxk0XWX0kdllllllllll.
         lxxxxxxk0XWMMMMNNNNX0xollllllc
        :xxxxOKNMMMMMMMMNNNNNNNNKOxolll:
       'xOKNMMMMMMMMMMMMNNNNNNNNNNNNKOdl,
       OWMMMMMMMMMMMMMMMNNNNNNNNNNNNNNNXk.
       .lkXMMMMMMMMMMMMMNNNNNNNNNNNNN0d:.
         ;;ckXMMMMMMMMMMNNNNNNNNNN0d:..
          ,;;;lkXMMMMMMMNNNNNNX0d:'...
           ,;;;;;cxXMMMMNNNN0d;......
            ,;;;;;;;lkXMNKd:........
             ,;;;;;;;,,c:''''......
              .

`,

	`              .'..................
             lx0NWWWWWWWWWWWWWWNOdc
            cxxxxOXWWWWWWWWWN0xllll:
           ;xxxxxxxkKNWWWWKkolllllll,
          .xxxxxxxxxxx0XOdlllllllllll.
         .dxxxxxxxxk0KNNKOxollllllllll.
         dxxxxxxOKNMMMMWNNNNKOdlllllllc
        lxxxk0XWMMMMMMMWNNNNNNNX0kdllll:
       ck0XWMMMMMMMMMMMWNNNNNNNNNNNX0kol'
      .KMMMMMMMMMMMMMMMWNNNNNNNNNNNNNNNKd
       .oONMMMMMMMMMMMMWNNNNNNNNNNNNNOo;
        .;;oONMMMMMMMMMWNNNNNNNNNNOo,'.
          ;;;:oONMMMMMMWNNNNNNXkl,'''.
           ,;;;;;oOWMMMWNNNXkl,'.'''.
            ,;;;;;;:d0WWXkl,.'.'..'.
             ';;;;;;,,;c,'''''''..


`,

	`              .'''''..............
             lx0NWWWWWWWWWWWWWWXOdc
            cdddx0NWWWWWWWWWXOxllll:
           :ddddddxOXWWWWN0xllllllll;
          ,dddddddddxOX0xolllllllllll'
         'ddddddddxOKNWX0kdlllllllllll.
        .ddddddk0XWMMMMNNNNX0kolllllllc.
       .dddxOKNMMMMMMMMNNNNNNNNKOxollll:
       dOKNMMMMMMMMMMMWNNNNNNNNNNNNKOxol'
      :NMMMMMMMMMMMMMMWNNNNNNNNNNNNNNNX0:
       'o0WMMMMMMMMMMMWNNNNNNNNNNNNNXko,
        .;:d0WMMMMMMMMMNNNNNNNNNNXkl,'.
         .;;;:xKWMMMMMWNNNNNNNKxc''''.
           ,;;;;:xXMMMWNNNNKd:''''''.
            ,;;;;;;lkNWN0d:''''''''.
             ';;;;;,,,c;''''''''''


`,

	`              .''''''.............
             lxKNWWWWWWWWWWWWWWXOdc
            cdddx0NWWWWWWWWWXOdllll:
           cddddddx0NWWWWXOdllllllll;
          ;dddddddddxKXkdllllllllllll,
         ,ddddddddk0XWNKOxollllllllllc'
        'dddddxOKNMMMMWWWWNKOxolllllllc.
       'ddxk0XWMMMMMMMWWWWWWWWN0kdollll:.
      .xOXWMMMMMMMMMMMWWWWWWWNNWNWX0kolc,
      dWMMMMMMMMMMMMMMWWWWWWNNWNWNNNNNKk.
       ,d0WMMMMMMMMMMMWWWWNNNNNNNNNNXkl.
        .;:xXMMMMMMMMMWWNWNNNNNNNKx:''.
         .,;;lkXMMMMMMWNNWNNNNOd:''''.
           ,;;;;lOWMMMWNNNNOl,''''''.
            ,;;;;;:dKMWXkl,''''''''.
             .;;;,,,,:c'''''''''''.
                            .....

`,

	`              ''''''''............
             lxKWWWWWWWWWWWWWWWXOdc
            cdddkKWWWWWWWWWNKkdoooo:
           cddddddkXWWWWNKkooooooooo;
          :dddddddddON0xooooooooooool,
         :dddddddxOKNWX0kdoooooooooooc,
        :dddddk0XWMMMWWWWWX0kooooooool:'
       ;ddxOKNMMMMMMMWWWWWWWWWXOxoooool:'
      ;k0NMMMMMMMMMMMWWWWWWWWWWWWNKOxoo:,
      kWMMMMMMMMMMMMMWWWWWWWWWWWWWWWWN0d
       ;xKMMMMMMMMMMMWWWWWWWWWWWWWWWKxc.
        .;ckNMMMMMMMMWWWWWWWWWWWN0o;''.
         .;;;o0WMMMMMWWWWWWWWXkl,''''.
           ,;;;:dKMMMWWWWWKxc'''''''.
            ';;;;;ckNWNOd:'''''''''.
             .;;;,,,,c,''''''''''''
                           ......

`,

	`              ''''''''............
             cxKWWWWWWWWWWWWWWWKOdc
            cdddkXWWWWWWWWWN0kdoooo:
           cddddddONWWWWXOxooooooool;
          cddddddddxXKkdooooooooooooc;
         cdddddddk0XWNKOxoooooooooooo:,
        cddddxOKWMMMMWWWWWKOdooooooool:,
       cddk0XWMMMMMMMWWWWWWWWNKkdoooooc:,
      lOKWMMMMMMMMMMMWWWWWWWWWWWWN0kool:;
     .0MMMMMMMMMMMMMMWWWWWWWWWWWWWWWWXOc.
       ;xXMMMMMMMMMMMWWWWWWWWWWWWWWN0d:.
        .;lOWMMMMMMMMWWWWWWWWWWWXkl,''.
         .,;:dKMMMMMMWWWWWWWW0xc,''''.
           ,;;;ckNMMMWWWWNOo;'''''''.
            ';;;;;dKMWKxc,''''''''''
             .,,,,,,::'''''''''''''
                           ......

`,

	`              ',''''''............
             cxKWWWWWWWWWWWWWWNKOdc
            cdddONWWWWWWWWWX0kdoooo:
           cdddddxKWWWWNKOdooooooool;
          cddddddddON0xoooooooooooooc;
         cddddddxOKWMN0kdoooooooooooo:;
        lddddk0NMMMMMMMMMN0koooooooooc::
       ldxOKWMMMMMMMMMMMMMMMWX0xoooooo:::
      d0XMMMMMMMMMMMMMMMMMMMMMMMWKOxooc::.
     'KMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN0x:.
       :kNMMMMMMMMMMMMMMMMMMMMMMMMMXOo:.
        .;oKMMMMMMMMMMMMMMMMMMMWKxc,,,'
         .,,ckNMMMMMMMMMMMMMNOo;,,,,,'
           ',,;oKMMMMMMMWKxc,,,,,,,,'
            .,,,,ckWMNOo;,,,,,,,,,,'
             .,,,,,,c,'''''',,,,,,'
                          .......

`,

	`              ',,''''''...........
             :dKWWWWWWWWWWWWWWNKOdc
            cddd0WWWWWWWWWNKOxooooo:
           cdddddkNWWWWX0kdooooooool:
          cdddddddxXKkdoooooooooooooc:
         lddddddk0NMWKOxoooooooooooolc:.
        lddddOKWMMMMMMMMWKOdooooooooocc:.
      .odk0XWMMMMMMMMMMMMMMMNKkdooooolcc:.
     .kKNMMMMMMMMMMMMMMMMMMMMMMMN0kdooccc,
     ;KMMMMMMMMMMMMMMMMMMMMMMMMMMMMWXOo:,
       cOWMMMMMMMMMMMMMMMMMMMMMMMMWKkl;,
        .;dXMMMMMMMMMMMMMMMMMMMNOo;,,,,
         .,,l0WMMMMMMMMMMMMWKxl,,,,,,,
           ',,:xNMMMMMMMNOo;,,,,,,,,,
            .,,,;oKMWKxc,,,,,,,,,,,,
             .,,,,,:;''',,,,,,,,,,'
                          .......

`,

	`              ',,,'''''...........
             ;dXWWWWWWWWWWWWWWNKOd:
            :dddKWWWWWWWWWNKOxooooo:
           cddddd0WWWWNKOxoooooooool:
          ldddddddOXOxdoooooooooooooc:
        .ldddddxOXWWX0kdoooooooooooolcc.
       .odddk0NMMMMMMMMWX0kooooooooooccc.
      .odOKNMMMMMMMMMMMMMMMWXOxoooooolccc.
     'OXWMMMMMMMMMMMMMMMMMMMMMMWKOdoolccc:
     :XMMMMMMMMMMMMMMMMMMMMMMMMMMMMX0xc:;
       cOWMMMMMMMMMMMMMMMMMMMMMMMMN0dc;,
        .:kNMMMMMMMMMMMMMMMMMMMXkl,,,;,
         .,;dKMMMMMMMMMMMMMNOd:,,,,,,,
           ',,cOWMMMMMMWKxc,,,,,,,,,,
            .,,,:kWMNOo;,,,,,,,,,,,,
              ',,,,:,',,,,,,,,,,,,,
                          .......

`,

	`              .,,,,''''...........
             ,dXWWWWWWWWWWWWWWN0kd:
            :oodXWWWWWWWWWX0kxoodoo:
           coooodXWWWWX0kddooddododlc
          loooooodX0kddddddodododooolc.
        .loooodkKNMNKOxddddddooooodollc.
       .ooodOKWMMMMMMMMNKOdoodooddodllll.
      'oxOXWMMMMMMMMMMMMMMMN0kdddoodollll,
     :0NMMMMMMMMMMMMMMMMMMMMMMWX0xddollllc
     cXMMMMMMMMMMMMMMMMMMMMMMMMMMMNKOoc:;.
       c0MMMMMMMMMMMMMMMMMMMMMMMMMXOo:;;.
        .:OWMMMMMMMMMMMMMMMMMMW0d:,,,;;
         .';xNMMMMMMMMMMMMWKxl;,,,,,;;
           .,,dXMMMMMMMXko;,,,,,,,,,;
            .,,,oKMW0d:,,,,,,,,,,,,,
              .'',:,,,,,,,,,,,,,,,,
                          .......

`,

	`              .;,,,''''..........
             'dXWWWWWWWWWWWWWWX0kx;
            ;oodNWWWWWWWWNX0kdddddo:
           cooooxWWWWNKOkddddddddddlc
          cooooooOXOxddddddddddddddolc.
        .looooxOXWWX0kdddddddddddddolll.
       .ooox0NMMMMMMMMWXOxddddddddddllll'
      'ok0NMMMMMMMMMMMMMMMNKOxddddddollll;
     lKNMMMMMMMMMMMMMMMMMMMMMMNKkdddollllc.
     :XMMMMMMMMMMMMMMMMMMMMMMMMMMWXOxlc:;.
       lKMMMMMMMMMMMMMMMMMMMMMMMMWKxl;;;.
        .c0WMMMMMMMMMMMMMMMMMMXkl;,,;;;.
          'cOWMMMMMMMMMMMMNOd:,,,,,,;;.
           .,;kWMMMMMMN0dc,,,,,,,,,;;
            .',;xNMXkl;,,,,,,,,,,,,;
              .'',;,,,,,,,,,,,,,,,,
                          .......

`,

	`              .;;,,,'''..........
             .oXWWWWWWWWWWWWWWX0kd,
            ,ooxWWWWWWWWWNKOkdddddo;
           :oooo0WWWWX0Oxdddddddddooc
          cooooodX0kxddddddddddddddool.
        .looodkKNMNKOxdddddddddddddoool.
       .oookKWWWWMMMMMN0kddddddddddooooo,
      ,dkKNWWWWWWMMMMMMMMWX0kdddddddooooo:
     oKWWWWWWWWWWMMMMMMMMMMMMWXOxdddoooool'
     :XWWWWWWWWWWMMMMMMMMMMMMMMMMN0kolc:;'
       lKWWWWWWWWMMMMMMMMMMMMMMMMNOd:;;;.
        .lKWWWWWWMMMMMMMMMMMMW0xc,,,;;;.
          'l0WWWWMMMMMMMMWKxl;,,,,,;;;.
           .'c0WWMMMMWKkl;,,,,,,,,,;;.
             ''c0MNOo:,,,,,,,,,,,,;;
              .'';,,,,,,,,,,,,,,,,,
                         ........

`,

	`              .;;,,,'''.........
             .lXWWWWWWWWWWWWWNX0kd'
            'ookWWWWWWWWWXKOkdddddo;
           :ooooNWWWNK0kdddddddddddoc
          coooooOKOkdddddddddddddddool.
        .looodOXWWX0kddddddddddddddoooo.
       .oodOXWWWWMMMMWKOxddddddddddooooo;
      ,dOXWWWWWWWMMMMMMMMNKOxdddddddoooooc
     dXWWWWWWWWWWMMMMMMMMMMMMX0kddddoooool;
     ;XWWWWWWWWWWMMMMMMMMMMMMMMMNKOdolc::,
       lKWWWWWWWWMMMMMMMMMMMMMMMWKkl;;;;'
        .oXWWWWWWMMMMMMMMMMMMXko;;,;;;;.
          .oXWWWWMMMMMMMMXOd:;;;;;;;;;.
           .'oXWWMMMMNOd:;;,;,;,;,;;;.
             .,dNWKxc;,;;,;,;,;,;,;;
              ..,,,,,,,,,,,;,;,;,;,
                         .......'

`,

	`               ;;;,,,''........
             .lXWWWWWWWWWWWWWNK0kd.
            .loOWWWWWWWWNX0Oxdddddd,
           ;ooodWWWWXKOxddddddddddddc
          :oooooX0kxdddddddddddddddddl.
        .loook0NMNKOxdddddddddddddddddo.
       'lox0NWWWMMMMMX0kdddddddddddddddd;
      ;x0XWWWWWWMMMMMMMMWX0kdddddddddddddc
     xXWWWWWWWWWMMMMMMMMMMMMNKOxdddddddddoc
     ;KWWWWWWWWWMMMMMMMMMMMMMMMWX0xdoolc:,
       cXWWWWWWWMMMMMMMMMMMMMMMMN0d:;;;;'
        .oNWWWWWMMMMMMMMMMMMW0xc;;;;;;;.
          .xNWWWMMMMMMMMWKxl;;;;;;;;;;.
            ,xNWMMMMWKxl;;;;;;;;;;;;;.
             .:OWXko;;;;;;;;;;;;;;;;
               .,'',,,,,,;;;;;;;;;,
                         ......''

`,

	`               ;;;,,,''........
              cXWWWWWWWWWWWWWNK0ko.
            .llKWWWWWWWWNK0Oxdddddd'
           ,lllOWWWNX0Oxdddddddddddd:
          :llllOKOkxdddddddddddddddddl.
        .clldOXWWX0kxdddddddddddddddddd.
       .lokKNWWWMMMMWKOxdddddddddddddddd;
      ;x0NWWWWWWMMMMMMMWNKOxdddddddddddddl
     kNWWWWWWWWWMMMMMMMMMMMWX0xdddddddddddl
     ,KWWWWWWWWWMMMMMMMMMMMMMMMN0kdddolc:;
       cXWWWWWWWMMMMMMMMMMMMMMMWKkl:;;;;,
        .dNWWWWWMMMMMMMMMMMMXko:;;;;;;;'
          'kNWWWMMMMMMMWXOd:;;;;;;;;;;.
            ;0WWMMMMXOo:;;;;;;;;;;;;;.
             .lXN0dc;;;;;;;;;;;;;;;,
               .''',,,,;;;;;;;;;;;,
                         ......',

`,

	`               ,;;;,,''.......
              :KWWWWWWWWWWWWWXKOkl
            .clXWWWWWWWNXK0kxxxxxxd.
           'lllXWWWXKOkxxxxxxxxxxxxx:
          ;llloX0Oxxxxxxxxxxxxxxxxxxxl
        .cllx0NMNKOxxxxxxxxxxxxxxxxxxxd.
       .ldOXNNNMMMMWX0kxxxxxxxxxxxxxxxxx;
      ;kKNNNNNNMMMMMMMMNK0xxxxxxxxxxxxxxxl
     kNNNNNNNNNMMMMMMMMMMMMNKkxxxxxxxxxxxdo
     '0NNNNNNNNMMMMMMMMMMMMMMMNKOxxxdolc:;.
       cXNNNNNNMMMMMMMMMMMMMMMMN0d:;;;;;,
        .dNNNNNMMMMMMMMMMMMN0xc;;;;;;;;'
          'ONNNMMMMMMMMN0xl;;;;;;;;;;;.
            cXNMMMMN0xc;;;;;;;;;;;;;;.
             .xWKkl;;;;;;;;;;;;;;;;,
               ..'',,,;;;;;;;;;;;;'
                         .....'',

`,

	`               ':;;,,''.......
              ;KWWWWWWWWWWWWNXKOxc
            .:lNWWWWWWWNXKOkxxxxxxd.
           .lldWWWNX0Okxxxxxxxxxxxxx;
          ,lllkKOkxxxxxxxxxxxxxxxxxxxl
        .:lokKWWX0kxxxxxxxxxxxxxxxxxxxd.
       .ld0NNNWMMMMNKOxxxxxxxxxxxxxxxxxx;
      ;kKNNNNNWMMMMMMMWX0kxxxxxxxxxxxxxxxl
     kNNNNNNNNWMMMMMMMMMMMWKOxxxxxxxxxxxxxo
     .ONNNNNNNWMMMMMMMMMMMMMMWX0kxxxdolcc:.
       :XNNNNNWMMMMMMMMMMMMMMMWKkl::::::;.
        .xNNNNWMMMMMMMMMMMWXko:;;;:::::,
          ,0NNNMMMMMMMWKko:;;;;;;;::::.
            oNNMMMWKko:;;;;;;;;;;;::;.
             .0XOo:;;;;;;;;;;;;;;;:;
               ..'',,,;;;;;;;;;;;:'
                         ....'',,

`,

	`               .::;,,''......
              '0WWWWWWWWWWWWNX0Ox:
             ;oNWWWWWWNXK0Okxxxxxko.
           .clOWWNXK0Oxxxxxxxxxxxkkx'
          'lloK0Okxxxxxxxxxxxxxxxkkkkc
         :ldOXMWX0kxxxxxxxxxxxxxxkkkkkd.
       .lxKNNNMMMMWX0kxxxxxxxxxxxkkkkkkk,
      ;OXNNNNNMMMMMMMMNKOkxxxxxxxkkkkkkkkl
     xNNNNNNNNMMMMMMMMMMMWX0kxxxxkkkkkkkkkd.
     .kNNNNNNNWMMMMMMMMMMMMMMNKkxkkxxdolcc.
       :XNNNNNWMMMMMMMMMMMMMMMXOocc:::::;.
        .kNNNNWMMMMMMMMMMMN0dc;;;::::::,
          :KNNWMMMMMMMXOdl;;;;;;;:::::.
           .kNWMMMNOdc;;;;;;;;;;;::::.
             :X0xc;;;;;;;;;;;;;;;::;
               ..'',,;;;;;;;;;;;;:'
                        .....',,;

`,

	`               .::;;,,'......
              .0WWWWWWWWWWWWNX0kx,
             ,dWWWWWWWNXK0Okxxxxxkl
           .clKWWNXKOkxxxxxxxxxxkkkx.
          .clxKOkxxxxxxxxxxxxxxxkkkkk;
         ;lxKNMNKOxxxxxxxxxxxxxxkkkkkko.
       .lkXXXNMMMMWKOxxxxxxxxxxxkkkkkkkk'
      ;OXXXXXNMMMMMMMWX0kxxxxxxxkkkkkkkkkl
     dXXXXXXXNMMMMMMMMMMMNKOxxxxkkkkkkkkkkd.
      xXXXXXXNMMMMMMMMMMMMMMNKOxkkkkxdoolc'
       :KXXXXNMMMMMMMMMMMMMMMWKxllcccccc:.
        .OXXXNMMMMMMMMMMMWKko:;;:cccccc;
          lXXNMMMMMMMN0xo:;;;;;;:ccccc'
           '0NMMMWKxl;;;;;;;;;;;:ccc:.
             xKko:;;;;;;;;;;;;;;:cc;
               ...',,;;;;;;;;;;;:c'
                        ....'',;;

`,

	`                ::;;,,'......
              .OWWWWWWWWWWWNXK0kd.
             'xWWWWWWWNXK0OkxxxxkO:
            :oNWNXK0OkxxxxxxxxxxkOOd.
          .cl00OkkxkxxxxxxxxxxxxOOOOk,
         ,okXWWX0kxxxxxxxxxxxxxxOOOOOOl
       .lOXXXWMMMWX0kkxxxxxxxxxxOOOOOOOk.
      ;OXXXXXWMMMMMMMNKOkxxxxxxkOOOOOOOOOc
     lXXXXXXXWMMMMMMMMMMWX0kxxxkOOOOOOOOOOd.
      oXXXXXXWMMMMMMMMMMMMMWX0kkOOOkxxdool,
       ;KXXXXWMMMMMMMMMMMMMMMXOoolccccccc.
        .0XXXWMMMMMMMMMMMXOdc;;:ccccccc;
          dXXWMMMMMMWXOdc;;;;;;:cccccc'
           :XWMMMXOo:;;;;;;;;;;;cccc:.
            '0Odc;;;;;;;;;;;;;;;ccc;
               ...',,;;;;;;;;;;;cc.
                        ....',,;:

`,

	`                :::;,,'.....
               xWWWWWWWWWWWNXKOko.
             .xWWWWWWNXK00OkkxxxOk,
            ;xNWNXK0OkkkkkkkkkxkOOOo
          .:dK0OkkkkkkkkkkkkkkkkOOOOk.
         'dONMNKOkkkkkkkkkkkkkkkOOOOOOc
        l0XXNMMMMWXOkkkkkkkkkkkOOOOOOOOx.
      ;OXXXXNMMMMMMMWX0kkkkkkkkOOOOOOOOOO:
     :XXXXXXXMMMMMMMMMMWNKOkkkkOOOOOOOOOOOd.
      lXXXXXXMMMMMMMMMMMMMMNKOkOOOOkkxddoo;
       ;KXXXXMMMMMMMMMMMMMMMN0doolllllllc.
        .0XXXMMMMMMMMMMMN0xl:;;llllllll;
         .kXXMMMMMMMN0xl:;;;;;;clllllc'
           dXMMMN0xc;;;;;;;;;;;cllll:.
            l0xl:;;;;;;;;;;;;;;:lll;
               ...',,;;;;;;;;;;:lc.
                        ....',;;

`,

	`                ,c:;,,'.....
               oNNNNNNNNNNNNXKOxl
             .kNNNNNNNXK0OOkkkkk0k.
            ,0NNXK0OOkkkkkkkkkkO000c
           :O0OkkkkkkkkkkkkkkkkO0000x.
         'x0WWX0kkkkkkkkkkkkkkkO00000O;
        l0KKWMMMMNKOkkkkkkkkkkk00000000d.
      ,OKKKKWMMMMMMMNKOkkkkkkkk000000000O;
     'KKKKKKWMMMMMMMMMMNKOkkkkO00000000000o
      :KKKKKWMMMMMMMMMMMMMNKOkO000OOkkxddo;
       ,KKKKWMMMMMMMMMMMMMMWKxddoollllllc.
        .0KKWMMMMMMMMMMWKOo:;;cllllllll:
         .OKWMMMMMMWKko:;;;;;;:lllllll'
          .ONMMWKko:;;;;;;;;;;:lllll:
            kko:;;;;;;;;;;;;;;;llll,
               ...',,;;;;;;;;;;clc.
                        ...',,;:

`,

	`                'c:;;,'.....
               cNNNNNNNNNNNNX0kx;
             .kNNNNNNXXK0OkkkkkO0d.
            'KNNXK0Okkkkkkkkkkk0000;
           cK0Okkkkkkkkkkkkkkkk00000d
         .xKMWX0kkkkkkkkkkkkkkO000000O,
        l0KXMMMMWX0kkkkkkkkkkkO00000000o
      'kKKKKMMMMMMMWX0kkkkkkkk0000000000O'
     .0KKKKKMMMMMMMMMMWX0kkkkk000000000000l
      '0KKKKMMMMMMMMMMMMMWX0kO00000OOkxxdd:
       '0KKKMMMMMMMMMMMMMMMXOdddooooooool.
        '0KKMMMMMMMMMMMXOxc;;:ooooooooo:
         '0KMMMMMMMXOdc;;;;;;;looooool.
          .0MMMXOdc;;;;;;;;;;;cooooo:
           .kdl;;;;;;;;;;;;;;;:oool'
               ...'',;;;;;;;;;:loc.
                        ...',;;c

`,

	`                .c::;,'....
               ;XNNNNNNNNNNXK0kx'
             .xNNNNNNXKK0Okkkkk0Ko.
            ,XNXKK0OkkkkkkkkkkkKKKO;
           oKOOkkkkkkkkkkkkkkkOKKKKKo
         .kNMNKOkkkkkkkkkkkkkk0KKKKKKk'
        c00NMMMMNKOkkkkkkkkkkkKKKKKKKKKl
      .k000NMMMMMMMNKOkkkkkkkOKKKKKKKKKKk.
      O0000NMMMMMMMMMMNKOkkkk0KKKKKKKKKKK0c
      .0000NMMMMMMMMMMMMMNKOkKKKK00OOOkkxx:
       .000NMMMMMMMMMMMMMMN0dxxddoooooool.
        '00NMMMMMMMMMMN0xl:::looooooooo;
         ,0XMMMMMMN0xl:::::::coooooool.
          ;XMMWKxl::::::::::::oooooo:
           cxo::::::::::::::::loooo'
                ...',;::::::::coo:
                        ...',;:l

`,

	`                 cc:;,'....
               ,KNNNNNNNNNNXKOxd.
             .xNNNNNNXK00OkkkkOK0c.
            cXNXK00Okkkkkkkkkk0KKKk,
          .k0OkkkkkkkkkkkkkkkkKKKKK0c.
         ,OMWX0OkkkkkkkkkkkkkOKKKKKKKx.
        l00MMMMWX0OkkOkkkkkkk0KKKKKKKK0:
      .x000MMMMMMMWX0OOkkkkkOKKKKKKKKKKKx.
      d0000WMMMMMMMMMWX0OOOkOKKKKKKKKKKKK0;
       k000WMMMMMMMMMMMMNKOO0KKKKK00OOkkkx:
       .O00WMMMMMMMMMMMMMWKxxxxdddooooool.
        '00WMMMMMMMMMWKkoc::coooooooooo;
         :0WMMMMMWKkoc:::::::ooooooool.
          oWMMXOo::::::::::::loooooo;
           doc::::::::::::::::ooool.
                ...',;::::::::loo;
                       ...',;::l

`,

	`                 :c:;,'....
               ,0NNNNNNNNNNX0Oxl
             .kNNNNNXXK00Okkkk0KO;.
            oNXKK0OOkkkkkkkkkkKKKKd'.
          ;0OOOOOOOOOOOkOkkkkOKXXXXO:.
         :KMWX0OOOOOOOOOOOOOOKXXXXXXKd.
        oOKMMMMNKOOOOOOOOOOOOKXXXXXXXX0;
      .xOOKMMMMMMMNKOOOOOOOO0XXXXXXXXXXKo
      :OOOKMMMMMMMMMWX0OOOOOKXXXXXXXXXXXXO.
       xOOKMMMMMMMMMMMMWX0O0XXXXKKK00OOOkk:
       .OOKMMMMMMMMMMMMMMNkxkxxxddddddddl.
        ,OKMMMMMMMMMMNOdc:::ddddddddddd,
         lKMMMMMMN0xl:::::::lddddddddl.
          kMMN0xl::::::::::::ddddddd,
          .oc::::::::::::::::lddddl.
                ...';;::::::::ddd,
                       ...',;:ll.

`,

	`                 ,c:;,'....
               'ONNNNNNNNNXK0kx:
             'kNNNNXXKK0OOOOOOKXk,.
           .kXXK00OOOOOOOOOOOOXXXKl'.
          o0OOOOOOOOOOOOOOOOOKXXXXXk,.
         oNMNKOOOOOOOOOOOOOOOXXXXXXXKl.
        oONMMMWX0OOOOOOOOOOOKXXXXXXXXXk,
       dOONMMMMMMWX0OOOOOOOOXXXXXXXXXXXKl
      .OOONMMMMMMMMMNKOOOOOKXXXXXXXXXXXXXk.
       lOONMMMMMMMMMMMWXKOOXXXXXXKKK00OOOk,
        kONMMMMMMMMMMMMMN0xkkkxxxdddddddl.
        ,OXMMMMMMMMMW0xl:::oddddddddddd,
         oXMMMMMWKko:::::::cdddddddddc.
         .XMWKko::::::::::::oddddddd'
          ;c;;:::::::::::::::dddddc
               ...',;::::::::lddo'
                       ...';::dc

`,

	`                 .cc:,'....
               'kNNNNNNNNNXKOkx'
             ,kNNNNXXK00OOOOOOXXd'.
           ,OXKK00OOOOOOOOOOOKXXX0:'.
         .O0OOOOOOOOOOOOOOOOOXXXXXXd,'
         xMWX0OOOOOOOOOOOOOOKXXXXXXX0:'
        dOMMMMWKOOOOOOOOOOOOXXXXXXXXXXd'
       lOOMMMMMMMNKOOOOOOOOKXXXXXXXXXXX0:
       kOOMMMMMMMMMWX0OOOO0XXXXXXXXXXXXXXo
       ;OOMMMMMMMMMMMMNKOOKXXXXXXXKKK000Ok.
        xOWMMMMMMMMMMMMWKxOkkkkxxxxxxxxxc
        ,OWMMMMMMMMWXko:::lxxxxxxxxxxxd.
         dWMMMMWXOdc:::::::dxxxxxxxxx:
         .WMN0dc:::::::::::cxxxxxxxo.
          c:,;::::::::::::::lxxxxx:
               ...',;::::::::dxxo.
                       ..',;:cd;

`,

	`                 .cc:;,...
               .dXNNNNNNNXXKOxd.
             ,kNNNXXKK00OOOOOKXKl..
           :0XK00OOOOOOOOOOOOXXXXk;,'.
         ,0OOOOOOOOOOOOOOOOOKXXXXXKc,'
        .0MWXOOOOOOOOOOOOOOOXXNXXXXXk;'
        o0MMMMN0OOOOOOOOOOOXXXXXXXXNNKl'
       :k0MMMMMMWX0OOOOOOO0XNXNNXNNNNNXk,
       lk0MMMMMMMMMNKOOOOOXXXXXXNNNNNNNNKc
       .k0MMMMMMMMMMMWXOO0XXXXXXXXXKKK000k.
        d0MMMMMMMMMMMMWKxkOOkkkkxxxxxxxx:
        'OMMMMMMMMMNOdc::cxxxxxxxxxxxxo.
         kMMMMMN0xl:::::::lxxxxxxxxxx;
         :MWKkl::::::::::::dxxxxxxxl.
          :,,:::::::::::::::xxxxxx,
               ...',;:::::::lxxxl.
                      ...',::od'

`,

	`                  :c:;,...
               .oXNNNNNNXXK0kkc.
             ;kNNXXKK00OOOOOOXNO:..
          .lKKK00OOOOOOOOOOOKNNNXd,,,.
         o0OOOOOOOOOOOOOOOOONNNNNNO:,,.
        .XMNKOOOOOOOOOOOOOOXNNNNNNNXo,,
        oXMMMWXOOOOOOOOOOO0NNNNNNNNNN0:'
       .kXMMMMMMNKOOOOOOOOXNNNNNNNNNNNXd'
       'kXMMMMMMMMWXOOOOOKNNNNNNNNNNNNNN0'
        kXMMMMMMMMMMMX0O0NNNNNNNNXXXKKKK0x
        lKMMMMMMMMMMMMNkkOOOOkkkkxxxxxxx;
        'KMMMMMMMMWKxl:::dxxxxxxxxxxxxo.
         KMMMMWKko:::::::cxxxxxxxxxxx,
         xMNOdc:::::::::::lxxxxxxxxc
         .,',::::::::::::::oxxxxxd'
               ...',::::::::dxxx:
                      ..',;::xo.

`,

	`                  ,c:;,...
               .lKXXXXXXXXKOxx,
             ;kXXXXKK00OOOOO0NNx;..
          .oKK00OOOOOOOOOOOONNNNKc,,,.
         O0OOOOOOOOOOOOOOOOKNNNNNNx,,,.
        .WWX0OOOOOOOOOOOOO0NNNNNNNNKc,,.
        lWMMMNKOOOOOOOOOOOXNNNNNNNNNNx;,
        xWMMMMMWXOOOOOOOOKNNNNNNNNNNNNKc.
        xWMMMMMMMMN0OOOO0NNNNNNNNNNNNNNNx.
        dNMMMMMMMMMMNKOONNNNNNNNNNXXXXKKKl
        :NMMMMMMMMMMMN0x00OOOOkkkkkkkkkx,
        'NMMMMMMMMXko:::dkkkkkkkkkkkkkl.
         NMMMMNOdc:::::::dkkkkkkkkkkd,
         KWKxl::::::::::::xkkkkkkkxc
         .'',:::::::::::::cxkkkkkd.
               ...,;:::::::cxkkx;
                      ..',;:lkl.
                            .
`,

	`                  'cc;,...
               .c0XXXXXXXKKOxd.
            .:kXXXKK00OOkkkkXNXo,..
          ,dKK0OOkkkkkkkkkk0NNNNO;,,,'
        .0OkkkkkkkkkOOOOOOONNNNNNXl,,,'
        ,MWKOkkOOOOOOOOOOOXNNNNNNNNk;,,.
        cMMMWX0OOOOOOOOOOKNNNNNNNNNNXl,,
        oMMMMMWNKOOOOOOO0NNNNNNNNNNNNNO:.
        oMMMMMMMMWKOOOOONNNNNNNNNNNNNNNXl
        cMMMMMMMMMMWXOOXNNNNNNNNNNNXXXXKK,
        ,MMMMMMMMMMMWKx0000OOOOkkkkkkkkd'
        .MMMMMMMMNOdc::lkkkkkkkkkkkkkkc
        .WMMMWKkl:::::::lkkkkkkkkkkko'
         WXOd::::::::::::lkkkkkkkkx:
         ..',:::::::::::::okkkkkko.
               ...,;:::::::dkkkx,
                      ..';::dk:
                            .
`,

	`                  .cc;,...
               .cOXXXXXXXK0kkl.
            .:kXXXKK00OOkkkONW0c,..
          ;xK00OOkkkkkkkkkONWWWNd;;;,'.
        ,0OkkkkkkkkkkkkkkkKWWWWWW0:;;;,
        cWNKOkkkkkkkkkkkk0WWWWWWWWNd;;;.
        cWWWWKOkkkkkkkkOONWWWWWWWWWW0:;,
        cWWWWWWXOkkOOOOONWWWWWWWWWWWWNd;.
        :WWWWWWWWX0OOOOXWWWWWWWWWWWWWWW0;
        ;WWWWWWWWWWX0OKWWWWWWWWWNNNNNXXX0.
        ,WWWWWWWWWWWXx000000OOOOOkkkkkko.
        'WWWWWWWW0xl::ckkkkkkkkkkkkkkx:
        .WWWWXOdc::::::ckkkkkkkkkkkkl'
        .WKxl:::::::::::cxkkkkkkkkx,.
         ..',::::::::::::ckkkkkkkc.
               ..',;::::::ckkkko.
                     ..',;:ckx,
                           .
`,

	`                  .;c:,'..
               .:kXXXXXXKK0xk,.
            .:kXXKK00OOkkkkKWWk;,...
         .:x00OOkkkkkkkkkk0WWWWXl;;;;,.
        c0kkkkkkkkkkkkkkkOWWWWWWWx;;;;;
        dWX0kkkkkkkkkkkkONWWWWWWWWKc;;;.
        oWWWN0kkkkkkkkkOXWWWWWWWWWWWx;;,
        oWWWWWN0OkkkkkkXWWWWWWWWWWWWWKc;
        lWWWWWWWNKOkkkKWWWWWWWWWWWWWWWWx.
        lWWWWWWWWWNKO0WWWWWWWWWWWWNNNNNXx
        cWWWWWWWWWWXkkKK00000OOOOOOOOOkl.
        :WWWWWWWKko:::xOOOOOOOOOOOOOOx,.
        :WWWN0xl:::::::xOOOOOOOOOOOkc..
        ;XOo::::::::::::dOOOOOOOOOo'.
         ..',::::::::::::dOOOOOOk:.
               ..',;::::::dOOOkl.
                     ..',::dOd.
                           .
`,

	`                   ,c:,'.
                ;xKXXXXKK0Oxx..
            .:xKKK00OOOkkkkNMNo;,...
         .ck00OkkkkkkkkkkkXMMMWO:;;;;,.
        xOkkkkkkkkkkkkkkkKMMMMMWXl;;;;;.
        0WKOkkkkkkkkkkkkKWMMMMMMMWk:;;;.
        OWWWKOkkkkkkkkkKMMMMMMMMWMWXl;;,
        OWWWWNKOkkkkkk0WMMMMMMMMWWMWWk:;
        kWWWWWWWXOkkkOWMMMMMMMMMWMWWWMXl.
        xWWWWWWWWNKOONMMMMMMMMMMWWWWNNNX;
        xWWWWWWWWWNOxKKKK00000OOOOOOOOk:
        dWWWWWWXOd:::oOOOOOOOOOOOOOOOo'.
        oWWWKko:::::::oOOOOOOOOOOOOx;..
        oKxl:::::::::::lOOOOOOOOOOc..
         ..',:::::::::::ckOOOOOOd,.
               ..',::::::ckOOOk:.
                     ..';:ckOl.
                           .
`,

	`                   'c:;'.
                ,dKXXXXKK0kkc.
            .:xKKK00OOkkkk0MM0c;,...
         'ck0OOkkkkkkkkkkOWMMMWd;;;;;,'
        0kkkkkkkkkkkkkkkOWMMMMMMO:;;;;;.
        NNKkkkkkkkkkkkkOWMMMMMMMMNo;;;;.
        XWNN0kkkkkkkkkOWMMMMMMMMMMMO:;;,
        XNWNNN0kkkkkkkNMMMMMMMMMMMMMNo;;
        KWWNNNWN0kkkkNMMMMMMMMMMMMMMMMO:
        KWNNNWNNNXOkXMMMMMMMMMMMMWWWWWW0
        0NWNWNWNWN0dKKKKKK000000OOOOOOx,
        ONWNNNN0xc::ckOOOOOOOOOOOOOOOc..
        ONWN0dc::::::ckOOOOOOOOOOOOo'..
        xOd::::::::::::xOOOOOOOOOk;..
         ..',:::::::::::xOOOOOOOl..
               ..',::::::oOOOOx,.
                     ..,;:oOO:
                          .'
`,

	`                   .:c;'.
                ,o0XXXKK00xk,.
            .:xKK00OOkkkkkXMWx:;,..
         'lkOOkkkkkkkkkkkXMMMMKc:::::,'.
       .0kkkkkkkkkkkkkkkXMMMMMMWd::::::.
        NX0kkkkkkkkkkkkXMMMMMMMMM0c::::'
        NNNXOkkkkkkkkkKMMMMMMMMMMMNo:::'
        NNNNNKOkkkkkkKMMMMMMMMMMMMMMO::,
        NNNNNNNKkkkk0MMMMMMMMMMMMMMMMNo,
        NNNNNNNNN0k0MMMMMMMMMMMMMMWWWWNo
        NNNNNNNNNKxOKKKKKKK0000000OOOOo.
        XNNNNNKkl:::xOOOOOOOOOOOOOOOx;..
        XNNKkl:::::::dOOOOOOOOOOOOOc...
        Okl:::::::::::oOOOOOOOOOOd'...
          .',::::::::::lOOOOOOOk:..
               ..';:::::ckOOOOo..
                    ..',::xOx'
                          .
`,

	`                   .;c;'.
                'lOKXKK00Oxd..
            .:x0K00OOkkkkOWMXo:;,..
        .,lkOkkkkkkkkkkkOWMMMWk::::::;'.
       .0kkkkkkkkkkkkkkOWMMMMMMKc::::::,
       .NXOkkkkkkkkkkkOWMMMMMMMMWx:::::'
       .NNNKkkkkkkkkkOWMMMMMMMMMMM0c:::'
       .NNNNN0kkkkkkOWMMMMMMMMMMMMMNd::'
       .NNNNNNXOkkkOWMMMMMMMMMMMMMMMMO:'
       .NNNNNNNNKkkWMMMMMMMMMMMMMMMWWWK,
        NNNNNNNNXkxKKKKKKKKK000000000Oc.
        NNNNNXOo:::oO000000000000000o'..
        NNXOdc::::::lO000000000000k;...
        Odc::::::::::ck000000000Ol....
          ..,;:::::::::x0000000x,..
               ..';:::::dO000Oc..
                    ..';:lO0o.
                          .
`,

	`                    'c:'.
                .cxKKKK00kkc..
            .:d000OOkkkkkKMMOc:;,..
        .,lkOkkxxxxxxkkkXMMMMNo::::::;'.
       ,OxxxxxxxxxxkkkkXMMMMMMWk:::::::;
       ,NKkxxxxxxkkkkkXMMMMMMMMMKl:::::;
       'NNX0kxxxkkkkkXMMMMMMMMMMMWd::::;
       'NNNNXOkkkkkkXMMMMMMMMMMMMMM0c::,
       .NNNNNNKkkkkXMMMMMMMMMMMMMMMMNo:,
       .NNNNNNNXOkXMMMMMMMMMMMMMMMMMWWx,
       .NNNNNNNXOdKKKKKKKKKKKK0000000x,'
       .NNNNN0xc::cO000000000000000Oc...
       .NNKxl:::::::x0000000000000d'....
        ko:::::::::::d0000000000k;....
          ..,;::::::::oO0000000o...
               ..';::::ck0000x;.
                    ..,;:d0Oc.
                         .,
`,

	`                    .::'.
                .:d0KK00OOO;..
            .:dO0OOkkxxxkNMWx:::,'..
        .;okkkxxxxxxxxxOWMMMM0c::::::;'.
       ;OxxxxxxxxxxxxxOWMMMMMMXl::::::::
       ;X0kxxxxxxxxxxOWMMMMMMMMWk:::::::
       ;XXXOxxxxxxxxOWMMMMMMMMMMMKl:::::
       ,XXXX0kxxxxx0WMMMMMMMMMMMMMNd:::;
       ,XXXXXXOxxx0MMMMMMMMMMMMMMMMMOc:;
       'XXXXXXX0k0MMMMMMMMMMMMMMMMMMMXc;
       'XXXXXXX0d0XXKKKKKKKKKKKKK0000o';
       .XXXXXkl:::x0000000000000000x,..,
       .XXOd:::::::o0000000000000Oc.....
       .dc::::::::::cO0000000000d'....
          ..,;::::::::k0000000O:....
               ..,;::::d00000o'.
                    ..,;cO0k,
                         .
`,

	`                    .;c,.
                .;o0K00OO0k,'.
            .:oO0OOkkxxx0WMKlcc:;'..
        .;okkkxxxxxxxxxKMMMMWdccccccc;,.
       :kxxxxxxxxxxxxxXMMMMMMMOccccccccc
       cX0xxxxxxxxxxkXMMMMMMMMMXocccccc:
       :XXKkxxxxxxxkNMMMMMMMMMMMWxccccc:
       ;XXXXOxxxxxkNMMMMMMMMMMMMMM0lccc;
       ;XXXXX0kxxkWMMMMMMMMMMMMMMMMNoc:;
       ,XXXXXXKkkWMMMMMMMMMMMMMMMMMMWk:;
       ,XXXXXXKxxXXXKKKKKKKKKKKKKKKKO:,;
       'XXXX0d:::l00000000000000000o...;
       'XKkl::::::ck0000000000000x,....'
       .o:::::::::::d00000000000c......
          ..,;:::::::o00000000x,....
               ..,;:::ck0000Oc..
                    .',:d00d.
                        .'
`,

	`                    .,c;.
                .;oOK00OOKo;'..
            .:okOOkkxxxkNMWkcccc;'..
        .;okkxxxxxxxxxkWMMMMKlccccccc:,.
       ckxxxxxxxxxxxxOWMMMMMMNdccccccccc.
       cKOxxxxxxxxxxOWMMMMMMMMWkccccccc:.
       cXX0xxxxxxxx0MMMMMMMMMMMMKlccccc:.
       :XXXKkxxxxxKMMMMMMMMMMMMMMNdccc:;.
       ;XXXXKOxxxKMMMMMMMMMMMMMMMMMOcc:;
       ;XXXXXKOxXMMMMMMMMMMMMMMMMMMMXl;;
       ,XXXXXKkdKXXXXKKKKKKKKKKKKKKKd,;;
       ,XXXKxc:::k0000000000000000O;..';
       'X0d:::::::d00000000000000o.....'
       .c::::::::::lO0000000000k,......
          ..';:::::::k00000000l.....
               ..,;:::o00000x,..
                    .';cO00c.
                        .;
`,

	`                     '::..
                .,lk00OOOKl:,..
            .;lxOkkxxxx0WMNocclc:,..
        .:okkxxxxxxxxxKMMMMWkcccccccc:,.
       ckxxxxxxxxxxxxKMMMMMMM0lccccccccc.
       cKkxxxxxxxxxxXMMMMMMMMMXoccccccc:.
       cKKOxxxxxxxkNMMMMMMMMMMMWxcccccc;.
       :KKKOxxxxxkWMMMMMMMMMMMMMM0lccc:;.
       :KKKK0xxxOWMMMMMMMMMMMMMMMMNocc;;.
       ;KKKKK0k0MMMMMMMMMMMMMMMMMMMWx:;;.
       ;KKKKKOoOXXXXXXXXKKKKKKKKKKK0c,;;.
       ,KKKOo:;:o0KKKKKKKKKKKKKKKKd'..,;.
       ,Kkl:;;;;;cOKKKKKKKKKKKKKO;.....'
       .:;;;;;;;;;:xKKKKKKKKKK0o.......
          ..';;;;;;;l0KKKKKKKO;.....
               ..,;;;ckKKKK0l...
                   ..';d0Kk,.
                        .
`,

	`                     .;:'.
                .'cxOOOOKkoc;'.
            .;lxOkkxxxxXMM0lccll:,..
        ':oxxxxxxxxxxkNMMMMXocccccccl:,.
       cxxxxxxxxxxxxkWMMMMMMWdcccccccccc.
       c0kxxxxxxxxxOWMMMMMMMMMOcccccccc:.
       cKKkxxxxxxx0MMMMMMMMMMMMKlcccccc;.
       :KKKkxxxxxXMMMMMMMMMMMMMMNdcccc:;.
       :KKKKkxxxNMMMMMMMMMMMMMMMMMkcc:;;.
       ;KKKKKOkNMMMMMMMMMMMMMMMMMMM0c;;;.
       ,KKKK0xdKXXXXXXXXXXXXKKKKKKKx,;;;.
       ,KK0xc;;cOKKKKKKKKKKKKKKKKO:...,;.
       '0xc;;;;;:d0KKKKKKKKKKKKKd'.....,.
       .,;;;;;;;;;lOKKKKKKKKKKO:........
           .';;;;;;:xKKKKKKKKd'.....
               ..,;;;o0KKKKO:...
                   ..,:kKKo..
                       .;
`,

	`                     .,:,.
                 ':okOOOXdol:,.
            .;cdkkxxxxOWMWxlllllc;'.
        ':oxxxxxxxxxx0MMMMMOlllllllllc,.
       cxxxxxxxxxxxxKMMMMMMMKllllllllllc'
       c0xxxxxxxxxxXMMMMMMMMMNollllllll:.
       :00xxxxxxxkNMMMMMMMMMMMWxllllllc;.
       :000xxxxxOWMMMMMMMMMMMMMM0llllc;;.
       ;0000xxx0MMMMMMMMMMMMMMMMMXolc:;;.
       ;00000xKMMMMMMMMMMMMMMMMMMMNd:;;;.
       ,0000ko0XXXXXXXXXXXXXXXXXXX0c,;;;.
       '00Ol;;:dKKKKKKKKKKKKKKKKKd'..';;.
       '0d:;;;;;cOKKKKKKKKKKKKKO:......,.
       .',;;;;;;;:xKKKKKKKKKKKd'........
           .';;;;;;l0KKKKKKK0c......
               ..,;;:xKKKKKd'....
                   ..,l0K0:..
                       .:
`,

	`                     .':,..
                 .;lxOkK0oooc,.
            .,cdkkxxxxKMMKolllloc;'.
        ':oxxxxxxxxxxNMMMMNdlllllllllc,.
       :xxxxxxxxxxxkWMMMMMMWxlllllllllll,
       :OxxxxxxxxxOWMMMMMMMMMOlllllllll:'
       ;0OxxxxxxxKMMMMMMMMMMMMKlllllllc:'
       ;00OxxxxxXMMMMMMMMMMMMMMNdllllc::'
       ,000kxxkWMMMMMMMMMMMMMMMMWxllc:::.
       ,0000kOWMMMMMMMMMMMMMMMMMMMOc::::.
       '000OoxKXXXXXXXXXXXXXXXXXXXx,;:::.
       '00x:;;cOKKKKKKKKKKKKKKKK0:...,;:.
       .Ol;;;;;:dKKKKKKKKKKKKKKd,.....';.
        .,;;;;;;;cOKKKKKKKKKK0c.........
           .',;;;;:xKKKKKKKKx,.......
               ..,;;lOKKKK0c.....
                   .';xKKx,..
                       ,
`,

	`                     .':;..
                .':okOkK0odoc;..
           .':lxkxxxxxKMMXollllol:,..
       .;ldxxxxxxxxxkNMMMMNdllllllllll:,.
       :xxxxxxxxxxxOWMMMMMMWklllllllllll,
       :Oxxxxxxxxx0MMMMMMMMMM0lllllllll:'
       ;0OxxxxxxxXMMMMMMMMMMMMNollllllc;'
       ;00kxxxxkNMMMMMMMMMMMMMMWxllllc;;.
       ,000kxx0WMMMMMMMMMMMMMMMMMOllc;;;.
       ,000OxKMMMMMMMMMMMMMMMMMMMMKl:;;;.
       '000klOXXXXXXXXXXXXXXXXXXXX0:;;;;.
       '00d:;;oKKKKKKKKKKKKKKKKKKo'..,;;.
       .Ol;;;;;ckKKKKKKKKKKKKKKk;.....';.
        ,;;;;;;;;o0KKKKKKKKKKKl'........
          ..';;;;;:kKKKKKKKKO;.......
              ..';;;o0KKKKKo'....
                   .':kKKk,..
                       ;
`,

	`                     .,:;'.
                .,cdkOkK0oddl:'.
           .,coxkxxxxxKMMXollllooc;'.
       .coxxxxxxxxxxkNMMMMWxllllllllllc;.
       :xxxxxxxxxxxOWMMMMMMMOlllllllllll,
       :OxxxxxxxxxKMMMMMMMMMMKollllllll:'
       ;OkxxxxxxkNMMMMMMMMMMMMWxllllllc;'
       ;OOkxxxxOWMMMMMMMMMMMMMMMOllllc;;.
       ,OOOxxxKMMMMMMMMMMMMMMMMMMXoll:;;.
       ,OO0OkNMMMMMMMMMMMMMMMMMMMMNd:;;;.
       'O0OxoKXXXXXXXXXXXXXXXXXXXXKl,;;;.
       'OOd;;:xKXXXXXXXXXXXXXXXXXk,..';;.
       .Ol;;;;;l0XXXXXXXXXXXXXX0c.....';.
       .:;;;;;;;:dKXXXXXXXXXXKd'.........
          ..,;;;;;cOXXXXXXXX0:........
              ..,;;:dKXXXXKd'....
                  ..,cOXXO;..
                      .;:
`,

	`                     .,c;'.
               ..;lxOOkK0oddoc,..
          .':ldkxxxxxxXMMXollllool:,'.
       'oxxxxxxxxxxxkWMMMMWxlllllllllllc'
       :xxxxxxxxxxx0WMMMMMMM0lllllllllll'
       ;OxxxxxxxxxXMMMMMMMMMMXdllllllll:'
       ;OkxxxxxxOWMMMMMMMMMMMMWkllllllc;'
       ,OOxxxxxKMMMMMMMMMMMMMMMMKollll:;.
       ,OOOxxkNMMMMMMMMMMMMMMMMMMNdll:;;.
       ,OOOkOWMMMMMMMMMMMMMMMMMMMMWOc;;;.
       'OOOoxXXXXXXXXXXXXXXXXXXXXXXk,;;;.
       .OOl;;cOXXXXXXXXXXXXXXXXXX0:..';;.
       .Ol;;;;;dKXXXXXXXXXXXXXXKo'....';.
       .c;;;;;;;:kXXXXXXXXXXXXk;.......'.
         ..';;;;;;o0XXXXXXXXKl'.......
              .',;;:xKXXXXXx,.....
                  ..,cOXX0:...
                      .:c
`,

	`                     .,c;,.
               .,:oxOOkK0oddoc;'.
         .';coxkxxxxxxXMMXolllllooc:,..
       ,dxxxxxxxxxxxOWMMMMWkllllllllllll,
       :xxxxxxxxxxxKMMMMMMMMKlllllllllll'
       ;kxxxxxxxxkNMMMMMMMMMMNxllllllll:'
       ;Oxxxxxxx0WMMMMMMMMMMMMM0llllllc;'
       ,OOxxxxxXMMMMMMMMMMMMMMMMXdllll:;.
       ,OOkxxOWMMMMMMMMMMMMMMMMMMWkllc;;.
       'OOOxKMMMMMMMMMMMMMMMMMMMMMMKo:;;.
       'OOkoONNNNNNNNNNNNNNNNNNNNNN0:;;;.
       .OOc;;oKXXXXXXXXXXXXXXXXXXKo'.';;.
       .Oc;;;;:xXXXXXXXXXXXXXXXXx,....';.
       .c;;;;;;;cOXXXXXXXXXXXX0:.......'.
         .',;;;;;;dKXXXXXXXXXd'........
             ..,;;;ckXXXXXXk;.....
                  .';l0XXKc...
                      .cl
`,

	`                    ..;c:,..
              ..;ldkOkkK0ddddl:,..
        ..,:lxkxxxxxxxXMMXdllllllolc;'..
       ;xxxxxxxxxxxxOWMMMMWOllllllllllll;
       :xxxxxxxxxxxKMMMMMMMMKollllllllll'
       ;kxxxxxxxxONMMMMMMMMMMWkllllllllc'
       ;OxxxxxxxKMMMMMMMMMMMMMMKollllll;'
       ,OkxxxxkNMMMMMMMMMMMMMMMMWxllll:;.
       ,OOxxxKMMMMMMMMMMMMMMMMMMMMKllc;;.
       'OOkkNMMMMMMMMMMMMMMMMMMMMMMNd:;;.
       'OOxdKNNNNNNNNNNNNNNNNNNNNNNXo,;;.
       .Ok:;:xXXXXXXXXXXXXXXXXXXXXx,..;;.
       .kc;;;;cOXXXXXXXXXXXXXXXXO:....';.
       .l;;;;;;;oKXXXXXXXXXXXXKl'......'.
        ..,;;;;;;:xXXXXXXXXXXx,........
             .',;;;cOXXXXXXO:......
                 ..,;oKXXKl'..
                      .ll
`,

	`                    .';c:,..
              .,:lxOOkkK0ddddoc;'.
        .,:ldxxxxxxxxxXMMXdllllllool:;'.
       ;xxxxxxxxxxxxOWMMMMWOllllllllllll;
       :xxxxxxxxxxxXMMMMMMMMXollllllllll'
       ;kxxxxxxxxOWMMMMMMMMMMWOllllllllc'
       ;kxxxxxxxXMMMMMMMMMMMMMMXdllllll:'
       ,kkxxxx0WMMMMMMMMMMMMMMMMWOllllc;.
       ,kkxxxXMMMMMMMMMMMMMMMMMMMMXdll:;.
       'OkkOWMMMMMMMMMMMMMMMMMMMMMMWOc;;.
       'kOokXNNNNNNNNNNNNNNNNNNNNNNNk;;;.
       .kx;;cOXXXXXXXXXXXXXXXXXXXX0:..,;.
       .k:;;;;o0XXXXXXXXXXXXXXXXKl'....;.
       .l;;;;;;:dXXXXXXXXXXXXXXd'......'.
        .';;;;;;;ckXXXXXXXXXXO;.........
            ..';;;;l0XXXXXX0l'.....
                 ..,;dKXXXd'...
                     ..oo
`,

	`                    .';c:;'.
             .';coxOOkkKOdddddl:,..
       .,;coxxxxxxxxxkXMMXdlllllllooc:,'.
       ;xxxxxxxxxxxx0WMMMMM0llllllllllll;
       ;xxxxxxxxxxkXMMMMMMMMNdllllllllll'
       ;xxxxxxxxx0WMMMMMMMMMMM0llllllllc'
       ,kxxxxxxkNMMMMMMMMMMMMMMNxllllll:'
       ,kxxxxxKWMMMMMMMMMMMMMMMMMKolllc;.
       ,kkxxkNMMMMMMMMMMMMMMMMMMMMWkll:;.
       'kkxKMMMMMMMMMMMMMMMMMMMMMMMMKo;;.
       'kko0NNNNNNNNNNNNNNNNNNNNNNNN0c;;.
       .ko;;oKNNNNNNNNNNNNNNNNNNNNKo'.,;.
       .k:;;;:xXNNNNNNNNNNNNNNNNXd,....;.
       .l;;;;;;:kXNNNNNNNNNNNNNk;......'.
        ',;;;;;;;l0NNNNNNNNNN0c.........
           ..',;;;;oKNNNNNNKo'......
                ..';:xXNNXx,...
                     .'dd
`,

	`                   ..,:lc;'..
            ..,:ldkOkkkKOdddddoc;,..
       .;coxxxxxxxxxxkXMMXdllllllllolc:,.
       ;xxxxxxxxxxxx0WMMMMM0llllllllllll;
       ;xxxxxxxxxxkNMMMMMMMMNxllllllllll,
       ;xxxxxxxxxKMMMMMMMMMMMMKolllllllc'
       ,kxxxxxxOWMMMMMMMMMMMMMMWkllllll:'
       ,kxxxxkXMMMMMMMMMMMMMMMMMMXdllll;.
       'kxxx0WMMMMMMMMMMMMMMMMMMMMM0llc;.
       'kkkNMMMMMMMMMMMMMMMMMMMMMMMMNx:;.
       'kxdXNNNNNNNNNNNNNNNNNNNNNNNNXd;;.
       .kl;:xXNNNNNNNNNNNNNNNNNNNNNx,.';.
       .x;;;;ckNNNNNNNNNNNNNNNNNNk;....,.
       .l;;;;;;c0NNNNNNNNNNNNNN0:......'.
        ,;;;;;;;;oKNNNNNNNNNNXo'........
           ..,;;;;:dXNNNNNNXd,.......
                ..,;:kXNNNk,....
                     .,xx,
`,

	`                   ..,:lc;,..
           ..,:coxOOkkk0Oddddddl:;'..
       .coxxxxxxxxxxxkXMMXdlllllllllllc;'
       ;xxxxxxxxxxxxKWMMMMM0olllllllllll;
       ;xxxxxxxxxxONMMMMMMMMWkllllllllll,
       ;xxxxxxxxxXMMMMMMMMMMMMXdlllllllc'
       ,xxxxxxx0WMMMMMMMMMMMMMMW0llllllc'
       ,kxxxxkNMMMMMMMMMMMMMMMMMMNxllll:.
       'kxxxKMMMMMMMMMMMMMMMMMMMMMMKoll;.
       'kxOWMMMMMMMMMMMMMMMMMMMMMMMMWOc;.
       .kdkNWWWWNWNWNWNWWNWNWNWNWNWNNk;;.
       .k:;cONNNNNNNNNNNNNNNNNNNNNN0:..;.
       .d;;;;l0NNNNNNNNNNNNNNNNNN0c....,.
       .c;;;;;;oKNNNNNNNNNNNNNNXl'.....'.
       .;;;;;;;;:dXNNNNNNNNNNNd'........
          ..',;;;;:xXNNNNNNNx;.......
               ..',;cONNNNO;....
                    ..;kk;
`,

	`                  ..';:lc:,'..
          ..';codkOkkkk0Oddddddoc:,'..
       'odxxxxxxxxxxxkXMMXxllllllllllllc,
       ;xxxxxxxxxxxxKMMMMMMKolllllllllll;
       ;xxxxxxxxxxOWMMMMMMMMWkllllllllll,
       ,xxxxxxxxkXMMMMMMMMMMMMXxllllllll'
       ,xxxxxxxKMMMMMMMMMMMMMMMMKolllllc'
       ,xxxxx0WMMMMMMMMMMMMMMMMMMWOllllc.
       'xxxkNMMMMMMMMMMMMMMMMMMMMMMNxll:.
       'xxKMMMMMMMMMMMMMMMMMMMMMMMMMMKo;.
       .xo0WWWWWWWWWWWWWWWWWWWWWWWWWWKc;.
       .x;;oKNNNNNNNNNNNNNNNNNNNNNNXl'.,.
       .o;;;;dKNNNNNNNNNNNNNNNNNNXo'...'.
       .c;;;;;:xXNNNNNNNNNNNNNNNd'.....'.
       .;;;;;;;;:kNNNNNNNNNNNNk;.........
         ...,;;;;;cONNNNNNNNO:........
               ..,;;cONNNN0:....
                    ..:OO:
`,

	`                  ..,;clc:;'..
         ..';:ldxkkkkkk0Odddddddlc:,'..
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll;
       ;xxxxxxxxxxxxKMMMMMMKolllllllllll;
       ;xxxxxxxxxx0WMMMMMMMMWOllllllllll,
       ,xxxxxxxxONMMMMMMMMMMMMNkllllllll'
       ,xxxxxxkXMMMMMMMMMMMMMMMMXdllllll'
       'xxxxxKWMMMMMMMMMMMMMMMMMMM0olllc.
       'xxxOWMMMMMMMMMMMMMMMMMMMMMMWOllc.
       'xkXMMMMMMMMMMMMMMMMMMMMMMMMMMNx:.
       .xdXWWWWWWWWWWWWWWWWWWWWWWWWWWNd;.
       .o;:xNWWWWWWWWWWWWWWWWWWWWWWNx,.,.
       .l;;;:kNWWWWWWWWWWWWWWWWWWNk;...'.
       .:;;;;;:kNWWWWWWWWWWWWWWNk;.......
       .;;;;;;;;cOWWWWWWWWWWWW0:.........
         ..';;;;;;l0NWWWWWWW0c.........
              ..',;;l0WWWWKl.....
                    ..c00c.
`,

	`                 ..',:col:;,'..
        ..';:loxkkkkkkk0Odddddddolc;,'..
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll:
       ;xxxxxxxxxxxxKMMMMMMKdlllllllllll;
       ;xxxxxxxxxx0WMMMMMMMMW0llllllllll,
       ,xxxxxxxxONMMMMMMMMMMMMWkllllllll,
       ,xxxxxxkNMMMMMMMMMMMMMMMMNxllllll'
       'xxxxkXMMMMMMMMMMMMMMMMMMMMXdllll'
       'xxxKWMMMMMMMMMMMMMMMMMMMMMMM0olc.
       .xOWMMMMMMMMMMMMMMMMMMMMMMMMMMWOc.
       .dOWWWWWWWWWWWWWWWWWWWWWWWWWWWWO:.
       .l;cOWWWWWWWWWWWWWWWWWWWWWWWWO:.'.
       .c;;;lOWWWWWWWWWWWWWWWWWWWW0c.....
       .:;;;;;l0WWWWWWWWWWWWWWWW0c.......
       .;;;;;;;;oKWWWWWWWWWWWWKl.........
        ..',;;;;;;oKWWWWWWWWKl'........
              ..,;;;oKWWWWXo'.....
                   ..'lKKl.
`,

	`                 ..';:colc:,'...
       ..,;:loxkkkkkkkk0Oddddddddolc;,'..
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll:
       ;xxxxxxxxxxxkXMMMMMMXdlllllllllll;
       ,xxxxxxxxxxKWMMMMMMMMM0olllllllll;
       ,xxxxxxxx0WMMMMMMMMMMMMWOolllllll,
       ,xxxxxxONMMMMMMMMMMMMMMMMWkllllll'
       'xxxxONMMMMMMMMMMMMMMMMMMMMNxllll'
       'xxkXMMMMMMMMMMMMMMMMMMMMMMMMXdll.
       .xKWMMMMMMMMMMMMMMMMMMMMMMMMMMMKo.
       .dKWWWWWWWWWWWWWWWWWWWWWWWWWWWWKl.
       .:;oKWWWWWWWWWWWWWWWWWWWWWWWWXl'..
       .:;;;oKWWWWWWWWWWWWWWWWWWWWXo'....
       .;;;;;;oXWWWWWWWWWWWWWWWWXo'......
       .;;;;;;;;dXWWWWWWWWWWWWNd'........
        .',;;;;;;:dXWWWWWWWWNd'.........
             ..',;;:dXWWWWNd'.....
                   ..'oXXd.
`,

	`                ..',;:lolc:;,'..
       .,;cloxkkkkkkkkk0Odddddddddolc;,'.
       ;xxxxxxxxxxxxxkXMMXxlllllllllllll:
       ;xxxxxxxxxxxkXMMMMMMXdlllllllllll;
       ,xxxxxxxxxxKWMMMMMMMMMKolllllllll;
       ,xxxxxxxxKWMMMMMMMMMMMMW0olllllll,
       'xxxxxx0WMMMMMMMMMMMMMMMMWOllllll,
       'xxxx0NMMMMMMMMMMMMMMMMMMMMWOllll,
       'xxONMMMMMMMMMMMMMMMMMMMMMMMMNkll'
       .kXMMMMMMMMMMMMMMMMMMMMMMMMMMMMXx'
       .xNMMMMMMMMMMMMMMMMMMMMMMMMMMMMNx.
       .;:xNMMMMMMMMMMMMMMMMMMMMWMWMWx,..
       .;;;:xNMMMMMMMMMMMMMMMWWWWWNx,....
       .;;;;;:xNMMMMMMMWMWWWWWWWWx,......
        ;;;;;;;:xNMWWWWWWWWWWWWx,........
        ',;;;;;;;:xNWWWWWWWWNx,.........
            ...,;;;:xNWWWWWx,......
                  ...,xNNx'.
`,

	`              ...',;:clollc:;,'...
       .:cldxkkkkkkkkkk0Oddddddddddolc:;'
       ;xxxxxxxxxxxxxOXMMXxlllllllllllll:
       ;xxxxxxxxxxxkXMMMMMMXxlllllllllll:
       ;xxxxxxxxxkKMMMMMMMMMMKdlllllllll;
       ,xxxxxxxxKWMMMMMMMMMMMMMKdlllllll;
       ,xxxxxxKWMMMMMMMMMMMMMMMMM0olllll;
       ,xxxx0WMMMMMMMMMMMMMMMMMMMMW0olll,
       'xx0WMMMMMMMMMMMMMMMMMMMMMMMMW0ol,
       'ONMMMMMMMMMMMMMMMMMMMMMMMMMMMMWO,
       .OWMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0'
       .;cOWMMMMMMMMMMMMMMMMMMMMMMMMMO:..
       .;;;cOWMMMMMMMMMMMMMMMMMMMMWO:....
        ;;;;;cOWMMMMMMMMMMMMMMMMMO;......
        ;;;;;;;cOWMMMMMMMMMMMMMO:.......
        ,;;;;;;;;cOWMMMMMMMMWO:.........
           ...';;;;ckWMMMMWO;.......
                  ..';kWWk,..
                       .`,

	`             ...',;;:cloolc:;,,'...
       ,lodxxxkkkkkkkkk0Oddddddddoooolc:,
       :xxxxxxxxxxxxxOXMMXklllllllllllll:
       :xxxxxxxxxxxkXMMMMMMNxlllllllllllc
       :xxxxxxxxxkXMMMMMMMMMMXxlllllllllc
       cxxxxxxxkXMMMMMMMMMMMMMMKxlllllllc
       cxxxxxkXMMMMMMMMMMMMMMMMMMXdlllllc
       cxxxkKWMMMMMMMMMMMMMMMMMMMMMKdlllc
       cxxKMMMMMMMMMMMMMMMMMMMMMMMMMMKdlc
       lKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0o
       cKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXl
       .;oKMMMMMMMMMMMMMMMMMMMMMMMMMMXl'.
       .;;;oKMMMMMMMMMMMMMMMMMMMMMMKl'...
       .;;;;;l0MMMMMMMMMMMMMMMMMMKc......
       .;;;;;;;l0MMMMMMMMMMMMMMKc........
        ;;;;;;;;;l0WMMMMMMMMM0c.........
           ..',;;;;cOWMMMMM0:.......
                 ...':OWWO:...
                       .`,

	`           ....',;;:clooollc:;;,'....
       ;dxxxxxxxxxkkkkk0Oddddddooooooool:
       cxxxxxxxxxxxxxOXMMXklllllllllllllc
       lxxxxxxxxxxxONMMMMMMNkllllllllllll
       lxxxxxxxxxkXMMMMMMMMMMXxllllllllll
       oxxxxxxxOXMMMMMMMMMMMMMMXxllllllll
       dxxxxxkXMMMMMMMMMMMMMMMMMMXxllllll
       xxxxOXMMMMMMMMMMMMMMMMMMMMMMXxllll
       xxkXMMMMMMMMMMMMMMMMMMMMMMMMMMNxll.
       kXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXx.
       xNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWx.
       ;:xNMMMMMMMMMMMMMMMMMMMMMMMMMMWx,.
       ';;:dXMMMMMMMMMMMMMMMMMMMMMMNd,...
       .;;;;;oXMMMMMMMMMMMMMMMMMMXo'.....
       .;;;;;;;oXMMMMMMMMMMMMMMXo'.......
        ;;;;;;;;;oKMMMMMMMMMMKl'.........
          ..',;;;;;l0MMMMMMKc........
                ...',c0WM0c...
                       .`,

	`         ....',,;;:cclodoollc::;,,'....
       cxxxxxxxxxxxxkkk0Odddddoooooollllc
       lxxxxxxxxxxxxxOXMMXkllllllllllllll
       oxxxxxxxxxxxONMMMMMMNkllllllllllll
       dxxxxxxxxxONMMMMMMMMMMNkllllllllll
       xxxxxxxxONMMMMMMMMMMMMMMNkllllllll.
       xxxxxxONMMMMMMMMMMMMMMMMMMNkllllll.
      .xxxx0NMMMMMMMMMMMMMMMMMMMMMMNOllll.
      .xx0NMMMMMMMMMMMMMMMMMMMMMMMMMMWOll'
      'ONMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWO;
      .0WMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0,
       ;cOWMMMMMMMMMMMMMMMMMMMMMMMMMMMO:.
       ;;;ckNMMMMMMMMMMMMMMMMMMMMMMWk;...
       ';;;;:xNMMMMMMMMMMMMMMMMMMWx,.....
       .;;;;;;:xNMMMMMMMMMMMMMMNx,.......
       .;;;;;;;;:dXMMMMMMMMMMXd'.........
         ..',;;;;;;oKMMMMMMXo'........
                ...',l0MMKl....
                       .`,

	`        ...',,;;::cclooddoollcc:;;,,'....
       cxxxxxxxxxxxxxkk0Oddddoooolllllllc
       oxxxxxxxxxxxxxOXMMXkllllllllllllll
       dxxxxxxxxxxxONMMMMMMNkllllllllllll
       xxxxxxxxxxONMMMMMMMMMMNkllllllllll.
      .xxxxxxxx0NMMMMMMMMMMMMMMNOolllllll.
      .xxxxxx0WMMMMMMMMMMMMMMMMMMWOolllll'
      ,xxxx0WMMMMMMMMMMMMMMMMMMMMMMW0olll;
      :xxKWMMMMMMMMMMMMMMMMMMMMMMMMMMW0ol:
      lKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0o
      cKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXo
      .;oKMMMMMMMMMMMMMMMMMMMMMMMMMMMMKl'.
       ;;;l0WMMMMMMMMMMMMMMMMMMMMMMM0c...
       ;;;;;cOWMMMMMMMMMMMMMMMMMMWO:.....
       ';;;;;;ckWMMMMMMMMMMMMMMWk;.......
       .;;;;;;;;:xNMMMMMMMMMMNx,.........
        ..',;;;;;;:dXMMMMMMNd'.........
               ...',;oKMMKo'....
                       '`,
}
