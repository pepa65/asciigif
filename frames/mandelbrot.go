package frames

func init() {
	FrameMap["mandelbrot"] = DefaultFrameType(_mandelbrot)
}

var _mandelbrot = []string{
	`MMMMMMMMMMO.     .MMMXKO..: .     . '.:cO, ;MMMMMWMMMMMMMMMMMMMM
MMMMMMMMMMW.     .Nd,.                     ;cONNkKMMMMMMMMMMMMMM
MMMMMMMMMMMXl;'  .                             . cNMMMMMMMMMMMMM
MMMMMMMMMMMNd.                                     :;MMMMMMMMMMM
MMMMMMMMXx:                                          ,x0MMMMMMMM
MMMNKOMN;                                              'xKMMMMMM
MMM0  '                                                  ,OMMMMM
MMMWO.                                                    .xMMMM
MW0d                                                        OMMM
MMl                                                          kMM
Mk.                                                           XM
M'                                                            .W
0                                                              d
l                                                              .
'
.
.
,
o                                                              '
K                                                              O
M:                                                            'M
MO.                                                           XM
MMc                                                         .0MM
MWN0                                                        KMMM
MMMWO.                                                    .OMMMM
MMM0  :.                                                 :0MMMMM
MMMWWXMW;                                              ,0XMMMMMM
MMMMMMMMXOo.                                        .c00MMMMMMMM
MMMMMMMMMMMWk;.                                   .o:MMMMMMMMMMM
MMMMMMMMMMM0;'.  ..                           .' dMMMMMMMMMMMMMM
MMMMMMMMMMX.      M0c,                     ;kXMWKXMMMMMMMMMMMMMM
MMMMMMMMMM0.     .MMMMNN.':..    .'.' ldN, ;MMMMMNMMMMMMMMMMMMM`,

	`MMMMMMMMMMMN.      kMXo.                     'KWWMMMMMMMMMMMMMMM
MMMMMMMMMMMMk.     ,.                          ..do.WMMMMMMMMMMM
MMMMMMMMMMMMMWk'                                    lOXkWMMMMMMM
MMMMMMMMMMMXc.                                         .KMMMMMMM
MMWMMMMMM0..                                             'cMMMMM
MMMMOc;0o                                                  'KMMM
MMMMX'.                                                     .xMM
MMMMNo.                                                       ,X
MWWlc                                                          .
MMM;
MMc
M0
Mo
M'
x

.
O
M'
Mx
MN.
MMd
MWM,
MWWxd                                                          '
MMMMWk.                                                      .cN
MMMMK.                                                      ,OMM
MMWM0xoNx.                                                 :NMMM
MMWMMMMMMX,'                                             ;dMMMMM
MMMMMMMMMMMNx,                                        . NMMMMMMM
MMMMMMMMMMMMMNk'                                    x0W0WMMMMMMM
MMMMMMMMMMMMd      c,.                        .';Ok'WMMMMMMMMMMM
MMMMMMMMMMMX       kMNk'                   . 'MMMMMMMMMMMMMMMMM`,

	`MMMMMMMMMMMMWO      o:.                           .:X:'WMMMMMMMM
MMMMMMMMMMMMMMXo:.                                    .0WMMMMMMM
MMMMMMMMMMMMMNOc.                                       .:,NMMMM
MMMMMMMMMMMdx                                              ,OkMM
MMMMMNNKWWx                                                  'O0
MMMMMx  ,.                                                     '
MMMMMWO'
MMMMM0,
MMMX,.
MMWX'
MMW'
MM0
MM;
kM
 x
 .
 .
 O
0M
MM:
MM0
MMW,
MMMW.
MMMX,,
MMMMMKc
MMMMMNo'
MMMMMk. c;                                                     ,
MMMMMMMWMMO                                                  'XX
MMMMMMMMMMM00.                                             cX0MM
MMMMMMMMMMMMMMXd;                                       'ocNMMMM
MMMMMMMMMMMMMMO;'.  .                               . .XMMMMMMMM
MMMMMMMMMMMMNk      dd'                           .lMl,MMMMMMMM`,

	`MMMMMMMMMMMMMMMWd...                                     .KMMMMM
MMMMMMMMMMMMMMMMM0,.                                      .lNxXM
MMMMMMMMMMMWXMkc.                                             dW
MMMMMMMMMMMM..                                                 .
MMMMMM0ocxN'
MMMMMMK.. .
MMMMMMMN:
MMMMN0O.
MMMMK'
MMMNx
MMM0
MMMo
MMM'
.kN
  d


  O
,0M
MMM'
MMMd
MMMX
MMMW0.
MMMMK.
MMMMMXX'
MMMMMMMX;
MMMMMMN   .
MMMMMMKxd0Wc
MMMMMMMMMMMM.;                                                 '
MMMMMMMMMMMWMM0d'                                           . kM
MMMMMMMMMMMMMMMMM0:.                                     .;xW0NM
MMMMMMMMMMMMMMMNl                                        .0MMMM`,

	`MMMMMMMMMMMMMMMMMMMK:                                        .cX
MMMMMMMMMMMMMMMMKk:
MMMMMMMMMMMMMOdo
MMMMMMMMMNMMK:
MMMMMMMN. .c
MMMMMMWM0:
MMMMMMMWk.
MMMMM0co
MMMMMMc
MMMMWd
MMMM0.
MMMM,
cNMW
  dK
   :


   o
 .OX
xMMM
MMMM:
MMMMK'
MMMMWx.
MMMMMWc
MMMMMKdk
MMMMMMMWO.
MMMMMMWMx,
MMMMMMMN'.'x.
MMMMMMMMMMMMNc
MMMMMMMMMMMMM0kk.
MMMMMMMMMMMMMMMMWKd                                            .
MMMMMMMMMMMMMMMMMMMW:                                        .d`,

	`MMMMMMMMMMMMMMMMMMMk.
MMMMMMMMMMMMMMWXWk.
MMMMMMMMMMMMMMN..
MMMMMMMMMxc:Xk
MMMMMMMMMo. .
MMMMMMMMMMX.
MMMMMMMMMd.
MMMMMMW,,
MMMMMMMO
MMMMMWl
MMMMMK.
WMMMM:
 ,KMM
   x0
    o

    .
    d
   0K
.cNMM.
MMMMMl
MMMMMK.
MMMMMWd.
MMMMMMMO
MMMMMMW;:.
MMMMMMMMMk.
MMMMMMMMMMK.
MMMMMMMMMc  '
MMMMMMMMMOxoW0.
MMMMMMMMMMMMMMM',
MMMMMMMMMMMMMMWWMK;
MMMMMMMMMMMMMMMMMMM0`,

	`MMMMMMMMMMMMMMMMMMMl.
MMMMMMMMMMMMMMMMW';
MMMMMMMMMMK0OXM0;.
MMMMMMMMMM0   c
MMMMMMMMMMWk,
MMMMMMMMMMWo
MMMMMMMWOxk
MMMMMMMMk.
MMMMMMM0c
MMMMMMMx
MMMMMMW.
dKWMMMO
  .OMM.
    oW
     o


     x
    kW
  ,XMM,
OWMMMMO
MMMMMMW.
MMMMMMMd
MMMMMMMNx
MMMMMMMWd.
MMMMMMMWX0X
MMMMMMMMMMWk
MMMMMMMMMMWo'
MMMMMMMMMMO  .k.
MMMMMMMMMMNWKWMKl.
MMMMMMMMMMMMMMMMW,o
MMMMMMMMMMMMMMMMMMMo`,

	`MMMMMMMMMMMMMMMMMMO :
MMMMMMMWMMMMMMMMMWx
MMMMMMMMMMMW''.l0,
MMMMMMMMMMMM:..
MMMMMMMMMMMMMk.
MMMMMMMMMMMM,
MMMMMMMMM0.;
MMMMMMMMMWx
MMMMMMMMXc
MMMMMMMM0.
MMMMMMMW:
''kXMMMX.
    oMMo
     'X;
      ..


      ,.
     ;N:
   .kMMx
;:0NMMMX.
MMMMMMMWc
MMMMMMMMK.
MMMMMMMMXo.
MMMMMMMMMWo
MMMMMMMMMK,l
MMMMMMMMMMMM:
MMMMMMMMMMMMMx.
MMMMMMMMMMMM'
MMMMMMMMMMMW:;,kK:
MMMMMMMWMMMMMMMMMWk
MMMMMMMMMMMMMMMMMMd.`,

	`MMMMMMMMMMMMWMMMMMMNO
MMMMMMMMMMMWM0d:,ON,
MMMMMMMMMMMMMX.   .
MMMMMMMMMMMMMMWk'
MMMMMMMMMMWMMMd.
MMMMMMMMMMWkxk
MMMMMMMMMMMk.
MMMMMMMMMMXO
MMMMMMMMMMo
MMMMMMMMMN;
oKWWMMMMMK
  .;OWMMM:
     .OMM
       oK
        c


        o
       kX
     ,KMM.
 .'oKMMMMl
kWMMMMMMMX
MMMMMMMMMNl
MMMMMMMMMMd.
MMMMMMMMMMWK.
MMMMMMMMMMMo..
MMMMMMMMMMW00K.
MMMMMMMMMMWMMM0.
MMMMMMMMMMMMMMXd'
MMMMMMMMMMMMMX   .'
MMMMMMMMMMMWW0OdcKWl.
MMMMMMMMMMMMNMMMMMMM`,

	`MMMMMMMMMMMMMMMkc,,0W:
MMMMMMMMMMMMMMMX   '.
MMMMMMMMMMMMMMMW0o'
MMMMMMMMMMMMMMMMXc
MMMMMMMMMMMMMXNN.
MMMMMMMMMMMMX...
MMMMMMMMMMMMMX.
MMMMMMMMMMMNl'
MMMMMMMMMMMM;
MMMMMMMMMMMx
;,c0NMMMMMM,
    .;OMMMN
       'XMk
         Oo
          '


         .;
        .Ko
       :WM0
    .lKMMMN
ocxNWMMMMMM:
MMMMMMMMMMMk
MMMMMMMMMMMN'
MMMMMMMMMMMNx,
MMMMMMMMMMMMMK.
MMMMMMMMMMMMX..,
MMMMMMMMMMMMMWMM'
MMMMMMMMMMMMMMMMWd
MMMMMMMMMMMMMMMWk;.
MMMMMMMMMMMMMMMK   ';
MMMMMMMMMMMMMMMkd:,XN`,

	`MMMMMMMMMMMMMMMMMM.   ,
MMMMMMMMMMMMMMMMMMO:,
MMMMMMMMMMMMMMMMMMMO
MMMMMMMMMMMMMMMMMM;
MMMMMMMMMMMMMMW;,c
MMMMMMMMMMMMMMMO.
MMMMMMMMMMMMMMWk
MMMMMMMMMMMMMX,
MMMMMMMMMMMMWk.
NWNXMMMMMMMMW:
   .;cNWMMMMK
       'xWMMo
         'KM.
           K
           ;

           .
           :
          .X
         ,WM'
      .;0WMMd
....ddMMMMMMO
MMWNMMMMMMMMN;
MMMMMMMMMMMMW0'
MMMMMMMMMMMMMX:.
MMMMMMMMMMMMMMMk
MMMMMMMMMMMMMMMd'
MMMMMMMMMMMMMMWocx
MMMMMMMMMMMMMMMMMMd.
MMMMMMMMMMMMMMMMMMMX'
MMMMMMMMMMMMMMMMMMx'.
MMMMMMMMMMMMMMMMMM'   `,

	`MMMMMMMMMMMMMMMMMMMMMdc'
MMMMMMMMMMMMMMMMMMMMMNo
MMMMMMMMMMMMMMMMMMMMO,
MMMMMMMMMMMMMMMMNko0.
MMMMMMMMMMMMMMMMWx'
MMMMMMMMMMMMMMMWMN,
MMMMMMMMMMMMMMMN;;
MMMMMMMMMMMMMMMMx
XMMMMMMMMMMMMMMO
 odkcXMNMMMMMMM:
      ..xNMMMMX.
         .cNMMd
           .NM:
            .X.
             ;


             :
            'N'
           'WM;
         'dWMMx
     .';0WMMMMN.
.kO0dMMMMMMMMMMo
WMMMMMMMMMMMMMMk.
MMMMMMMMMMMMMMMMd
MMMMMMMMMMMMMMMNoo
MMMMMMMMMMMMMMMMMW:
MMMMMMMMMMMMMMMMNc'.
MMMMMMMMMMMMMMMMWKkK'
MMMMMMMMMMMMMMMMMMMM0;
MMMMMMMMMMMMMMMMMMMMMWO.
MMMMMMMMMMMMMMMMMMMMMl,`,

	`MMMMMMMMMMMMMMMMMMMMMMMMW:
MMMMMMMMMMMMMMMMMMMMMMMO,
MMMMMMMMMMMMMMMMMMMXxkK.
MMMMMMMMMMMMMMMMMMM0..
MMMMMMMMMMMMMMMMMMMM0.
MMMMMMMMMMMMMMMMMMKO.
MMMMMMMMMMMMMMMMMM0'
MMMMMMMMMMMMMMMMMNo
W:kMMWNMMMMMMMMMMN'
' ..''.oOKMMMMMMMl
          ,kNMMMM,
            'OWMK
              cMO
               lo
                .


                '
               xd
              xMO
            ;0MMN.
         .c0WMMMM,
; .;:,'OKNMMMMMMMo
Wo0MMMWMMMMMMMMMMW.
MMMMMMMMMMMMMMMMMWk.
MMMMMMMMMMMMMMMMMMk'
MMMMMMMMMMMMMMMMMMNK.
MMMMMMMMMMMMMMMMMMMMO'
MMMMMMMMMMMMMMMMMMMO  .
MMMMMMMMMMMMMMMMMMMW00W'
MMMMMMMMMMMMMMMMMMMMMMM0:
MMMMMMMMMMMMMMMMMMMMMMMMN`,

	`MMMMMMMMMMMMMMMMMMMMMMMMMMO
MMMMMMMMMMMMMMMMMMMMMMNOkW'
MMMMMMMMMMMMMMMMMMMMMMd. .
MMMMMMMMMMMMMMMMMMMMMMMM;
MMMMMMMMMMMMMMMMMMMWMXN'
MMMMMMMMMMMMMMMMMMMWWc.
MMMMMMMMMMMMMMMMMMMMW0
MMN0MMMMMMMMMMMMMMMMx.
KO, Ox0OdMMMMMMMMMMM:
         .,c0NMMMMMN.
             .dNMMM0
               .kMMl
                 cM;
                  x'
                   '

                   .
                  .'
                  O'
                 xM:
               ,0MMc
             ,OMMMMK
        .'cdNWMMMMMW.
WXo X0NXxMMMMMMMMMMMo
MMMNMMMMMMMMMMMMMMMMk.
MMMMMMMMMMMMMMMMMMMMMK.
MMMMMMMMMMMMMMMMMMMWW:'
MMMMMMMMMMMMMMMMMMMMMWM:
MMMMMMMMMMMMMMMMMMMMMMMN;
MMMMMMMMMMMMMMMMMMMMMMx  .
MMMMMMMMMMMMMMMMMMMMMMMXKM:
MMMMMMMMMMMMMMMMMMMMMMMMMMK`,

	`MMMMMMMMMMMMMMMMMMMMMMMMWKo;ck
MMMMMMMMMMMMMMMMMMMMMMMMMW;
MMMMMMMMMMMMMMMMMMMMMMMMMMMo
MMMMMMMMMMMMMMMMMMMMMMMMMWX
MMMMMMMMMMMMMMMMMMMMMMMMO,.
MMMMMMMMMMMMMMMMMMMMMMMMMo
MMMMMMMMMMMMMMMMMMMMMMMK'
MMMM',MWWWOWMMMMMMMMMMM0'
c.''  . .. lKdMMMMMMMMMo
              ':KNMMMMW.
                 .KMMM0
                   oMMl
                    oM'
                     O.
                     .


                     .
                     0.
                    xM;
                  .xMMo
                 ;NMMMX
              ;oNMMMMMW'
x;::  ..,' xWOMMMMMMMMMk
MMMM;:MMMMKMMMMMMMMMMMMX,
MMMMMMMMMMMMMMMMMMMMMMMK:
MMMMMMMMMMMMMMMMMMMMMMMMWc
MMMMMMMMMMMMMMMMMMMMMMMMO',
MMMMMMMMMMMMMMMMMMMMMMMMMMW
MMMMMMMMMMMMMMMMMMMMMMMMMMNl
MMMMMMMMMMMMMMMMMMMMMMMMMX.  .
MMMMMMMMMMMMMMMMMMMMMMMMNKkcl`,

	`MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMO,.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMO
MMMMMMMMMMMMMMMMMMMMMMMMMMMMWKO
MMMMMMMMMMMMMMMMMMMMMMMMMMMN;
MMMMMMMMMMMMMMMMMMMMMMMMMMMMX,
MMMMMMMMMMMMMMMMMMMMMMMMMMMk,
MMMMMMo;NMMMMMMMMMMMMMMMMMMX
MWKooo. ,;:o'cKWXMMMMMMMMMW;
'              . xKNMMMMMMM.
                   ;xMMMMM0
                     ,XMMM:
                       dMM'
                        xM.
                         0
                         '

                         .
                         ;
                         X
                        OM.
                      .OMM;
                     cNMMMd
                   l0MMMMMK
:'             ,.0WWMMMMMMM.
MMWkOx. clcx:dWMMMMMMMMMMMNo
MMMMMMkoWMMMMMMMMMMMMMMMMMMX.
MMMMMMMMMMMMMMMMMMMMMMMMMMMOd
MMMMMMMMMMMMMMMMMMMMMMMMMMMMX,
MMMMMMMMMMMMMMMMMMMMMMMMMMMW;..
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMK.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMk'`,

	`MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWW00d
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX;.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNc,
MMMMMMMMX:oMMMMMMMMMMMMMMMMMMMMMk
:NMMMkKkc  dkxO:;WMMMMMMMMMMMMM0.
 :o,.            .':xMMMMMMMMMNx
                     .,KNMMMMMW;
                        .xMMMMN.
                          ,XMMk
                           .NMc
                            'W'
                             x.



                             .
                             O.
                            :M'
                           ;WMc
                          cWMM0
                        ,OMMMMN.
                     ,:WWMMMMMW;
 cOc'           .,co0MMMMMMMMMMk
dWMMMKWKx  0K0XxcMMMMMMMMMMMMMMk'
MMMMMMMMNdkMMMMMMMMMMMMMMMMMMMMMk
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWo:
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX'.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWMNNd
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0`,

	`MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW,  .
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWo.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWx
MMMMMMMMMMMWMMMMMMMMMMMMMMMMMMMMMMMMMo'
NMMMMMMMMMM0,cWMMMMWWMMMMMMMMMMMMMMMMW.
 .,0MMMXxKdc  co:xk'kMMMMMMWMMMMMMMMX'
   :kc.             ..;:oMMWMMMMMMMM0.
o                        .;xMMMMMMMMd
                            .ONMMMMW,
                              'kMMMX
                                oMM0
                                 cMx
                                  xl
                                   c
                                   '


                                   '
                                  .c
                                  kd
                                 oWO
                                xMMO
                              ;KMMMX
.                           ;KWMMMMW'
l. .                     ,c0MMMMMMMMo
   cKx,.            .;oddMMWMMMMMMMMK,
.;cXMMMX0W0l  xkd0K:kMMMMMMMMMMMMMMWK;
MMMMMMMMMMMKcdMMMMMWMMMMMMMMMMMMMMMMMW'
MMMMMMMMMMMWMMMMMMMMMMMMMMMMMMMMMMMMMo'
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMx
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX:.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN'  `,

	`MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXo
MMMMMMMMMMMMMMMMMMMMMMMWMMMMMMMMMMMMMMMMMMMNOx'
MMMMMMMMMMMMMMMMKXMMMMMMMMMMMMMMMMMMMMMMMMMWd.
MMkNWWMMMMMMMMM;  0MMMMMKNMMMMMMMMMMMMMMMMWNX
M0.  'XMMXKcdlc'  ,,,:c, lXWWXMMMMMMMMMMMMM0.
Mo    oo'                  .' NMMMMMMMMMMMMk
MM0,                           ;.NMMMMMMMMNd
Kd'                               ,XWMMMMMN.
.                                   oWMMMMN
                                     .OMMM0
                                       kMM:
                                        0M'
                                         N.
                                         l




                                         l.
                                        .W.
                                        0M'
                                      .KMMc
                                     '0MMM0
.                                   xWMMMMX
NO:                              ':WMMMMMMN.
MWk,                          .o;MMMMMMMMMWd
Mc    dk:.                .;:.MMMMMMMMMMMMWx
MN...:XMMWWoOko,  :c:dd: dMMMMMMMMMMMMMMMMW0'
MMXWMMMMMMMMMMM:..KMMMMMWWMMMMMMMMMMMMMMMMWW0
MMMMMMMMMMMMMMMMNNMMMMMMMMMMMMMMMMMMMMMMMMMWl.
MMMMMMMMMMMMMMMMMMMMMMMWMMMMMMMMMMMMMMMMMMMMKO.
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN`,

	`MWMMMMMMMMMMMMMMMMMMMKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMd
MMMMMWMMMMMMMMMMMMMN: ;XMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0
MMMMK.c:cKMMMMMOWXKO   xOkk0Wd.xMMMMWMMMMMMMMMMMMMMNc.
MMMN'     MWk::                'doKkkMMMMMMMMMMMMMMMO
MMMMo.   .,                         'ON0MMMMMMMMMMMK.
MMMMMX:                               . OWNMMMMMMMMx
MMNk.                                     ;NMMMMMMWc
0Wc                                        .OMMMMMM'
 .                                           ,NMMMW
                                              .KMMK
                                               .XMO
                                                .Wx
                                                 o:
                                                  ,




                                                  ;
                                                 xc
                                                ;Mx
                                               .WMO
                                              ,XMMX
 ,                                           cMMMMM
XMl                                        '0MMMMMW.
MMW0;.                                   .oMMMMMMMWl
MMMMMX;                               . KMMMMMMMMMMk
MMMWc    .c.                        'KWXMMMMMMMMMMMX'
MMMN,    .MMKdx ...         .  ,OOW00MMMMMMMMMMMMMMMk
MMMMX'dldNMMMMMKMWW0   kX0kWMx,xMMMMWMMMMMMMMMMMMMMXc,
MMMMMMMMMMMMMMMMMMMNd.oNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMK
MNMMMMMMMMMMMMMMMMMMMNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMl`,

	`MMMMMMMMK.   oWMMMMMddWOkdc  'c:c'dkM  cMMMMMMMMMMMMMMMMMMMMMMXc
MMMMMMM0.     dMMKdo. .             .  ;odNXxXMMMMMMMMMMMMMMMMMW
MMMMMMMW;     c;.                            oKMWMMMMMMMMMMMMM0;
MMMMMMMMMKo.                                   '.0MMMMMMMMMMMMN,
MMMMMMMMx,.                                       :cMMMMMMMMMNd
MMMMMW.'                                            ,xMMMMMMMWd
MNxdX.                                                ;XMMMMMW,
MN:                                                    .xMMMM0
MWo                                                      dNMMk
d'                                                        cWWl
c                                                          xM'
                                                           .M.
                                                            x.
                                                            '




                                                            ;
                                                            0.
                                                           'M'
d                                                          OM'
o:                                                        dMMc
MMk                                                      xMMMx
MX' .                                                  '0MMMMX
MNOOW:                                               .oWMMMMMW;
MMMMMM;;                                            cOMMMMMMMMo
MMMMMMMW0c'                                      .ldMMMMMMMMMMd
MMMMMMMMMk:.                                  .:'XMMMMMMMMMMMMWo
MMMMMMMX'     oc,                         .. xWMMMMMMMMMMMMMMM0:
MMMMMMMK.     OMMWOk. '             '  :kkMW0NMMMMMMMMMMMMMMMMMN
MMMMMMMMN.   dNMMMMMk0M0N0k  ,dld,00M  :MMMMMMMMMMMMMMMMMMMMMMN`,
}
