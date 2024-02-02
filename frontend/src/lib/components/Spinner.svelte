<script lang="ts">
    export let size:string = "l"
    export let thickness:string = "l"
    export let color:string = "primary"
    export let isBlock:boolean = true

    const sizeVariants = {
        s: 'w-3 h-3',
        m: 'w-7 h-7',
        l: 'w-12 h-12'
    }

    const thicknessVariants = {
        s: 'border-2',
        m: 'border-5',
        l: 'border-8'
    }

    const colorVariants = {
        primary: 'border-primary',
        secondary: 'border-secondary',
        grey: 'border-gray-300'
    }

    let sizeClass = sizeVariants[size] || sizeVariants['l']
    let bThicknessClass = thicknessVariants[thickness] || thicknessVariants['l']
    let bColorClass = colorVariants[color] || colorVariants['primary']

    $: blockOrInline = isBlock ? "flex h-full w-full items-center justify-center" : "inline-block"
</script>

<div class="{blockOrInline}">
    <div class="{sizeClass}
                {bThicknessClass}
                {bColorClass}
                rounded-full
                animate-spin
                border-solid
                border-t-transparent"></div>
</div>

