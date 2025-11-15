export const addClassList = (element: Element, className: string) => {
    Object.values(element.children).forEach((e) => {
        e.classList.add(className)
        if (e.children.length > 0) addClassList(e, className)
    })
}