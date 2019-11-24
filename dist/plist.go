package dist

type Plist struct {
	BuildMachineOSBuild         string        `plist:"BuildMachineOSBuild"`
	BundleDevelopmentRegion     string        `plist:"CFBundleDevelopmentRegion"`
	BundleDocumentTypes         []interface{} `plist:"CFBundleDocumentTypes"`
	BundleExecutable            string        `plist:"CFBundleExecutable"`
	BundleIconFile              string        `plist:"CFBundleIconFile"`
	BundleIdentifier            string        `plist:"CFBundleIdentifier"`
	BundleInfoDictionaryVersion string        `plist:"CFBundleInfoDictionaryVersion"`
	BundleName                  string        `plist:"CFBundleName"`
	BundlePackageType           string        `plist:"CFBundlePackageType"`
	BundleShortVersionString    string        `plist:"CFBundleShortVersionString"`
	BundleSignature             string        `plist:"CFBundleSignature"`
	BundleSupportedPlatforms    []string      `plist:"CFBundleSupportedPlatforms"`
	Compiler                    string        `plist:"DTCompiler"`
	PlatformBuild               string        `plist:"DTPlatformBuild"`
	PlatformVersion             string        `plist:"DTPlatformVersion"`
	SDKBuild                    string        `plist:"DTSDKBuild"`
	SDKName                     string        `plist:"DTSDKName"`
	Xcode                       string        `plist:"DTXcode"`
	XcodeBuild                  string        `plist:"DTXcodeBuild"`
	ApplicationCategoryType     string        `plist:"LSApplicationCategoryType"`
	HighResolutionCapable       bool          `plist:"NSHighResolutionCapable"`
	HumanReadableCopyright      string        `plist:"NSHumanReadableCopyright"`
	PrincipalClass              string        `plist:"NSPrincipalClass"`

	// We want to remove exported type declarations from the plist
	// so we do not attempt to decode / encode it
	//ExportedTypeDeclarations    []interface{} `plist:"UTExportedTypeDeclarations"`
}

/* Sample love plist file
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>BuildMachineOSBuild</key>
	<string>15C50</string>
	<key>CFBundleDevelopmentRegion</key>
	<string>English</string>
	<key>CFBundleDocumentTypes</key>
	<array>
		<dict>
			<key>CFBundleTypeIconFile</key>
			<string>GameIcon</string>
			<key>CFBundleTypeName</key>
			<string>LÖVE Project</string>
			<key>CFBundleTypeRole</key>
			<string>Viewer</string>
			<key>LSHandlerRank</key>
			<string>Owner</string>
			<key>LSItemContentTypes</key>
			<array>
				<string>org.love2d.love-game</string>
			</array>
		</dict>
		<dict>
			<key>CFBundleTypeName</key>
			<string>Folder</string>
			<key>CFBundleTypeOSTypes</key>
			<array>
				<string>fold</string>
			</array>
			<key>CFBundleTypeRole</key>
			<string>Viewer</string>
			<key>LSHandlerRank</key>
			<string>None</string>
		</dict>
		<dict>
			<key>CFBundleTypeIconFile</key>
			<string>Document</string>
			<key>CFBundleTypeName</key>
			<string>Document</string>
			<key>CFBundleTypeOSTypes</key>
			<array>
				<string>****</string>
			</array>
			<key>CFBundleTypeRole</key>
			<string>Editor</string>
		</dict>
	</array>
	<key>CFBundleExecutable</key>
	<string>love</string>
	<key>CFBundleIconFile</key>
	<string>OS X AppIcon</string>
	<key>CFBundleIdentifier</key>
	<string>org.love2d.love</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>LÖVE</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>0.10.1</string>
	<key>CFBundleSignature</key>
	<string>LoVe</string>
	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>MacOSX</string>
	</array>
	<key>DTCompiler</key>
	<string>com.apple.compilers.llvm.clang.1_0</string>
	<key>DTPlatformBuild</key>
	<string>7C68</string>
	<key>DTPlatformVersion</key>
	<string>GM</string>
	<key>DTSDKBuild</key>
	<string>15C43</string>
	<key>DTSDKName</key>
	<string>macosx10.11</string>
	<key>DTXcode</key>
	<string>0720</string>
	<key>DTXcodeBuild</key>
	<string>7C68</string>
	<key>LSApplicationCategoryType</key>
	<string>public.app-category.games</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>NSHumanReadableCopyright</key>
	<string>© 2006-2016 LÖVE Development Team</string>
	<key>NSPrincipalClass</key>
	<string>NSApplication</string>
	<key>UTExportedTypeDeclarations</key>
	<array>
		<dict>
			<key>UTTypeConformsTo</key>
			<array>
				<string>com.pkware.zip-archive</string>
			</array>
			<key>UTTypeDescription</key>
			<string>LÖVE Project</string>
			<key>UTTypeIconFile</key>
			<string>GameIcon</string>
			<key>UTTypeIdentifier</key>
			<string>org.love2d.love-game</string>
			<key>UTTypeReferenceURL</key>
			<string>http://love2d.org/wiki/Game_Distribution</string>
			<key>UTTypeTagSpecification</key>
			<dict>
				<key>com.apple.ostype</key>
				<string>LOVE</string>
				<key>public.filename-extension</key>
				<array>
					<string>love</string>
				</array>
				<key>public.mime-type</key>
				<string>application/x-love-game</string>
			</dict>
		</dict>
	</array>
</dict>
</plist>*/
