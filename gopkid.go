package gopkid

var m map[string][]string
var b bool

// GetInfo ...
func GetInfo() map[string][]string {
	initInfo()
	return m
}

func initInfo() {
	if !b {
		b = true
		m = make(map[string][]string)

		m["爱加密"] = []string{"libexecmain", "libexec", "ijiami.dat"}
		m["爱加密企业版"] = []string{"ijiami.ajm"}
		m["梆梆加固"] = []string{"libsecmain", "libSecshell", "libSecShell-x86", "libsecexe"}
		m["梆梆加固企业版"] = []string{"libDexHelper", "libDexHelper-x86", "libdexjni"}
		m["腾讯乐固"] = []string{"libtup", "libshell", "libtxRes64", "libshella", "liblegudb", "libshellx", "mix.dex", "mixz.dex"}
		m["腾讯御安全"] = []string{"libtosprotection.armeabi", "libtosprotection.armeabi-v7a", "libtosprotection.x86", "libdev.so"}
		m["360加固"] = []string{"libjiagu_ls", "libprotectClass", "libprotectClass_x86", "libjiagu_art", "libjiagu_x86", "libjiagu"}
		m["腾讯御安全"] = []string{"libTmsdk-", "libtosprotection.armeabi-v7a", "libtosprotection.armeabi", "libtosprotection.x86"}
		m["百度加固"] = []string{"libbaiduprotect", "ibbaiduprotect_x86", "baiduprotect"}
		m["娜迦加固"] = []string{"libddog", "libfdog", "libchaosvmp", "libvdog", "libhdog"}
		m["娜迦加固企业版"] = []string{"libedog"}
		m["网秦加固"] = []string{"libnqshieldx86", "libnqshield"}
		m["阿里聚安全"] = []string{"libmobisec", "libmobisecx", "libdemolish", "libfakejni", "libzuma", "libzumadata", "libpreverify1", "aliprotect.dat", "libsgmain", "libsgsecuritybody", "libmobisec.so"}
		m["通付盾加固"] = []string{"libegis", "libegis-x86.so", "libNSaferOnly"}
		m["网易易盾"] = []string{"libnesec", "libsecert"}
		m["盛大加固"] = []string{"libapssec"}
		m["几维安全"] = []string{"libkwscmm", "libkwscr", "libkwslinker", "kdpdata", "libkdp"}
		m["灵御珊瑚"] = []string{"libreincp", "libreincp_x86", "rein_classes.jar"}
		m["APKProtect"] = []string{"libAPKProtect"}
		m["乐变加固"] = []string{"libzmvmpcls", "libkxqpzmvmpinf"}
		m["顶像科技"] = []string{"libx3g"}
		m["海云安加固"] = []string{"libitsec"}
		m["瑞星加固"] = []string{"librsprotect"}
		m["uu安全"] = []string{"libuusafe", "libuusafeempty"}
		m["DxShield"] = []string{"libdxbase"}
		m["DexProtector"] = []string{"dp.arm", "dp.arm-v7"}
		m["Kiro"] = []string{"libkiroro"}
		m["Medusah"] = []string{"libmd"}
		m["AppGuard"] = []string{"libAppGuard", "libAppGuard-x86"}
		m["Aproov"] = []string{"libapproov"}
		m["Kony"] = []string{"libkonyjsvm"}
		m["PangXie"] = []string{"libnsecure"}
		m["Firehash"] = []string{"libaurorabridge"}
		m["DexLoad"] = []string{"libdexload"}
		m["中国移动mogo"] = []string{"libcmvmp", "libmogosec_dex", "libmogosec_sodecrypt", "libmogosecurity"}

	}

}
