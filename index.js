import * as common from "./common.js"
import * as constants from './constants.js';

// ====================get element by id====================
let elementList = {};
constants.IdList.map( id => {
	elementList[id] = document.getElementById(id);
});
// ====================get element by id====================

const defaultContainerWidth = elementList.mContainer.offsetWidth;
const currentURL = new URL(window.location.href).origin + new URL(window.location.href).pathname;
const currentURLParams = new URLSearchParams(new URL(window.location.href).search);
let codeTextFromURL = currentURLParams.get("code");

init();

async function init() {
	fillCodeLanguageOptions();
	await setConfigData()
	hightlightCodeText();
}

// TODO: 驗證從 URL回傳回來的數字
async function setConfigData() {
	elementList.codeLangSelectpicker.value = 	currentURLParams.get("codeLang") || 
												localStorage.getItem("codeLang") || 
												constants.DefaultCodeLang;
	$('#codeLangSelectpicker').selectpicker('refresh');

	elementList.cssStyleSelector.value = 		currentURLParams.get("cssStyle") || 
												localStorage.getItem("cssStyle") || 
												constants.DefaultCssStyle;

	elementList.backgroundColorInput.value = 	currentURLParams.get("backgroundColor") || 
												localStorage.getItem("backgroundColor") || 
												constants.DefaultBackgroundColor;

	elementList.containerColorInput.value = 	currentURLParams.get("containerColor") || 
												localStorage.getItem("containerColor") || 
												constants.DefaultContainerColor;

	elementList.containerWidthInput.value = 	common.PixelStrToInt(currentURLParams.get("containerWidth")) ||
	 											localStorage.getItem("containerWidth") || 
												defaultContainerWidth;
	
	elementList.fontSizeInput.value = 			common.PixelStrToInt(currentURLParams.get("fontSize")) ||
												localStorage.getItem("fontSize") || 
												constants.DefaultFontSize;
	
	elementList.mCodeText.value = localStorage.getItem("codeText") || constants.DefaultCodeText;
	if (codeTextFromURL !== null) {
		try {
			const res = await getCodeContent(codeTextFromURL);
			const result = await res.json();
			if (res.ok) {
				// TODO: 移除行號
				let mRegex = /(<span class=\"token[a-zA-Z0-9-_ ]*\">)|(<\/span>)/ig
				elementList.mCodeText.value = result.message.replaceAll(mRegex, "");
				let eraseLineNumber = [];
				let mLineNumber = /[0-9]+\t/ig
				elementList.mCodeText.value.split("\n").map(it => {
					eraseLineNumber.push(it.replace(mLineNumber, ""))
				});
				elementList.mCodeText.value = eraseLineNumber.join("\n");
			} else {
				codeTextFromURL = null;
			}
		} catch(error) {
			console.log('error', error);
		}
	}
}

function fillCodeLanguageOptions() {
	let options = "";
	constants.SupportLanguage.map(it => {
		options += "<option>" + it + "</option>\n";
	});
	elementList.codeLangSelectpicker.innerHTML = options;
}

function hightlightCodeText() {
	elementList.mBackground.style.backgroundColor = elementList.backgroundColorInput.value;
	elementList.mContainer.style.backgroundColor = elementList.containerColorInput.value;
	elementList.mCodeStyle.href = constants.CodeStyleMap[elementList.cssStyleSelector.value];
	elementList.mCode.style.fontSize = elementList.fontSizeInput.value.toString() + "px";
	elementList.mContainer.style.width = elementList.containerWidthInput.value.toString() + "px";
	elementList.mCode.innerHTML = elementList.mCodeText.value.replaceAll("<", "&lt;");
	elementList.mCode.className = "language-" + elementList.codeLangSelectpicker.value;
	Prism.highlightElement(elementList.mCode);
	elementList.mPre.className = "";
	elementList.mCode.className = "";

	// add line number for last render
	let counterOfLines = 1;
	let addCodeLineNumberText = [];
	elementList.mCode.innerHTML.split("\n").map(it => {
		addCodeLineNumberText.push("<span class=\"token comment\">" + counterOfLines.toString() + "</span>\t" + it);
		counterOfLines++;
	});
	elementList.mCode.innerHTML = addCodeLineNumberText.join("\n");
}

// ====================windows on resize====================
// window.addEventListener("resize", () => {
// 	const newContainerWidth = parseInt(0.52356 * window.screen.availWidth + 36);
// 	elementList.containerWidthInput.value = newContainerWidth;
// 	elementList.mContainer.style.width = newContainerWidth + "px";
// });
// ====================windows on resize====================

// ====================color picker, on color change====================
elementList.backgroundColorInput.addEventListener("input", (e) => {
	localStorage.setItem("backgroundColor", e.target.value);
	elementList.mBackground.style.backgroundColor = e.target.value;
});
elementList.containerColorInput.addEventListener("input", (e) => {
	localStorage.setItem("containerColor", e.target.value);
	elementList.mContainer.style.backgroundColor = e.target.value;
});
// ====================color picker, on color change====================

// ====================css style selector, on option change====================
elementList.cssStyleSelector.addEventListener("change", (e) => {
	localStorage.setItem("cssStyle", e.target.value);
	elementList.mCodeStyle.href = constants.CodeStyleMap[e.target.value];
});
// ====================css style selector, on option change====================

// ====================font size input, on value change====================
elementList.fontSizeInput.addEventListener("change", (e) => {
	let val = parseInt(e.target.value);
	const min = parseInt(e.target.min);
	const max = parseInt(e.target.max);
	val = common.Clamp(val, min, max);
	e.target.value = val;
	localStorage.setItem("fontSize", val);
	elementList.mCode.style.fontSize = val.toString() + "px";
});
// ====================font size input, on value change====================

// ====================container width input, on value change====================
elementList.containerWidthInput.addEventListener("change", (e) => {
	let val = parseInt(e.target.value);
	const min = parseInt(e.target.min);
	const max = parseInt(e.target.max);
	val = common.Clamp(val, min, max);
	e.target.value = val;
	localStorage.setItem("containerWidth", val);
	elementList.mContainer.style.width = val.toString() + "px";
});
// ====================container width input, on value change====================

// ====================code language select picker, on value change====================
elementList.codeLangSelectpicker.addEventListener("change", (e) => {
	localStorage.setItem("codeLang", e.target.value);
	hightlightCodeText();
});
// ====================code language select picker, on value change====================

// ====================code textarea, event====================
elementList.mCodeText.addEventListener("input", (e) => {
	localStorage.setItem("codeText", e.target.value);
});
elementList.mCodeText.addEventListener("keydown", (e) => {
	// ref: https://stackoverflow.com/questions/6637341/use-tab-to-indent-in-textarea
	if (e.key == 'Tab') {
		e.preventDefault();
		let start = e.target.selectionStart;
		let end = e.target.selectionEnd;
		e.target.value = e.target.value.substring(0, start) + '\t' + e.target.value.substring(end);
		e.target.selectionStart = e.target.selectionEnd = start + 1;
	}
});
// ====================code textarea, event====================

// ====================edit preview button, on click====================
elementList.editPreviewBtn.addEventListener("click", () => {
	hightlightCodeText();
	elementList.editSection.className = common.ToggleVisable(elementList.editSection.className);
	elementList.previewSection.className = common.ToggleVisable(elementList.previewSection.className);
});
// ====================edit preview button, on click====================

// ====================clean catch button, on click====================
elementList.cleanCatchBtn.addEventListener("click", async () => {
	localStorage.clear();
	await setConfigData();
	hightlightCodeText();
});
// ====================clean catch button, on click====================

// ====================click for copy====================
[...document.getElementsByClassName('mCopyButton')].map(it => {
	it.addEventListener("click", () => {
		let beforeCopy = it.children[0];
		let afterCopy = it.children[1];
		beforeCopy.className = common.HideElement(beforeCopy.className);
		afterCopy.className = common.ShowElement(afterCopy.className);
		let copyValue = elementList[it.getAttribute('mTarget')].innerHTML;
		navigator.clipboard.writeText(copyValue);
		setTimeout(()=>{
			beforeCopy.className = common.ShowElement(beforeCopy.className);
			afterCopy.className = common.HideElement(afterCopy.className);
		}, 500);
	})
});
// ====================click for copy====================

// ====================generate image button, on click====================
elementList.generateImageBtn.addEventListener("click", async () => {
	hightlightCodeText();
	elementList.modalLoadingSpinner.className = common.ShowElement(
		elementList.modalLoadingSpinner.className
	);
	elementList.modalResultSuccess.className = common.HideElement(
		elementList.modalResultSuccess.className
	);
	elementList.modalResultFailed.className = common.HideElement(
		elementList.modalResultFailed.className
	);

	const uploadData = {
		"code_language": "language-" + elementList.codeLangSelectpicker.value,
		"code_content": elementList.mCode.innerHTML,
	};

	try{
		const res = await uploadCode(uploadData);
		const result = await res.json();
		if (res.ok) {
			elementList.modalResultSuccess.className = common.ShowElement(
				elementList.modalResultSuccess.className
			);

			let params = new URLSearchParams(new URL(constants.RequestHost).search);
			params.set("code", result.message);
			params.set("backgroundColor", elementList.backgroundColorInput.value);
			params.set("containerColor", elementList.containerColorInput.value);
			params.set("containerWidth", elementList.containerWidthInput.value.toString() + "px");
			params.set("fontSize", elementList.fontSizeInput.value.toString() + "px");
			params.set("cssStyle", elementList.cssStyleSelector.value);
			params.set("codeLang", elementList.codeLangSelectpicker.value);
			let linkURL = currentURL + "?" + params.toString();
			params.set("file", "image.png");
			let imgURL = constants.RequestHost + "/raw/code/image?" + params.toString();

			elementList.resultURLText.innerHTML = imgURL;
			elementList.resultBBCodeText.innerHTML = "[url=" + linkURL + "][img=" + imgURL + "][/url]";
		} else {
			elementList.modalResultFailed.className = common.ShowElement(
				elementList.modalResultFailed.className
			);
			elementList.resutlErrorText.innerHTML = result.message;
		}
	}catch(error){
		elementList.modalResultFailed.className = common.ShowElement(
			elementList.modalResultFailed.className
		);
		elementList.resutlErrorText.innerHTML = error;
		console.log(error);
	}

	elementList.modalLoadingSpinner.className = common.HideElement(
		elementList.modalLoadingSpinner.className
	);
	$("#resultModal").modal("handleUpdate");
});
// ====================generate image button, on click====================

async function uploadCode(uploadData) {
	let mHeaders = new Headers();
	mHeaders.append("Content-Type", "text/plain");
	let requestOptions = {
		method: 'POST',
		headers: mHeaders,
		body: JSON.stringify(uploadData),
		redirect: 'follow',
	};
	return fetch(constants.RequestHost+"/code", requestOptions);
}

async function getCodeContent(hashName) {
	return fetch(constants.RequestHost+"/code?code="+hashName);
}
