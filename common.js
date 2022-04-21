
function Clamp (val, min, max) {
    val = val < min ? min : val;
    val = val > max ? max : val;
    return val;
}

function ToggleVisable(classNameStr) {
    if (typeof(classNameStr) !== "string")
        return classNameStr;

    if (classNameStr.includes("visually-hidden")) {
        return classNameStr.replace("visually-hidden", "");
    }
    return classNameStr + " visually-hidden";
}

function HideElement(classNameStr) {
    if (!classNameStr.includes("visually-hidden")) {
        return classNameStr + " visually-hidden";
    }
    return classNameStr;
}

function ShowElement(classNameStr) {
    if (classNameStr.includes("visually-hidden")) {
        return classNameStr.replace("visually-hidden", "");
    }
    return classNameStr;
}

function PixelStrToInt(pixelStr) {
    if (pixelStr === null || pixelStr === undefined)
        return null;
    if (typeof(pixelStr) !== "string")
        return pixelStr;
    pixelStr.replace("px", "");
    return parseInt(pixelStr);
}

export {Clamp, ToggleVisable, HideElement, ShowElement, PixelStrToInt};