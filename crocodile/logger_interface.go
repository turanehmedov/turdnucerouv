/ *
  * Bu fayl Timsah Oyun Botunun bir hissəsidir.
  * Müəllif hüquqları (C) 2019 Viktor
  *
  * Bu proqram pulsuz bir proqramdır: onu yenidən paylaya və / və ya dəyişdirə bilərsiniz
  * tərəfindən nəşr olunduğu GNU Ümumi Dövlət Lisenziyası şərtlərinə görə
  * Pulsuz Proqram Vəqfi, ya Lisenziyanın 3-cü versiyası, ya da
  * (seçiminizə görə) hər hansı sonrakı versiya.
  *
  * Bu proqram faydalı olacağı ümidi ilə paylanmışdır,
  * lakin heç bir zəmanət olmadan;  hətta nəzərdə tutulan zəmanət olmadan
  * BELƏKƏLİ MƏQSƏD ÜÇÜN SATILIRILMA VƏ FİTNES.  Baxın
  * Daha ətraflı məlumat üçün GNU Ümumi İctimai Lisenziyası.
  *
  * GNU Ümumi İctimai Lisenziyasının bir nüsxəsini almalı idiniz
  * bu proqramla birlikdə.  Əks təqdirdə, <http://www.gnu.org/licenses/>.
  * /

package crocodile

type Logger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Traceln(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}
