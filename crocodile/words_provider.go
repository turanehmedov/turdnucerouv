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

 paket timsah

 idxal (
 "io"
 "io / ioutil"
 "riyaziyyat / rand"
 "simlər"
 "vaxt"
 )

 // WordsProviderReader oxucudan məzmun alır, sətrə çevirir, "\ n" ilə bölünür və təsadüfi sözü qaytarır
 WordsProviderReader struct {yazın
 wordsList [] simli
 }

 // NewWordsProviderReader WordsProviderReader yeni nümunəsini qaytarır
 func NewWordsProviderReader (r io.Reader) (* WordsProviderReader, səhv) {
 məzmun, səhv: = ioutil.ReadAll (r)
 səhv varsa! = sıfır {
 qayıt sıfır, səhv
 }

 contentString: = strings.TrimSpace (string (content))
 contentString = strings.ReplaceAll (contentString, "ё", "е")

 qayıt & WordsProviderReader {
 wordsList: strings.Split (contentString, "\ n"),
 }, sıfır
 }

 // GetWord təsadüfi sözü qaytarır
 func (w * WordsProviderReader) GetWord () (string, error) {
 rand.Seed (time.Now (). UnixNano ())
 indeks: = rand.Intn (len (w.wordsList))
 qayıtmaq üçün simlər.TrimSpace (w.wordsList [index]), sıfır
 }