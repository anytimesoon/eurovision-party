<script lang="ts">
    interface Props {
        size?: string;
        thickness?: string;
        color?: string;
        isBlock?: boolean;
    }

    let {
        size = "l",
        thickness = "l",
        color = "primary",
        isBlock = true
    }: Props = $props();

    const sizeVariants = {
        s: 'w-3 h-3',
        sm: 'w-[20px] h-[20px]',
        m: 'w-7 h-7',
        l: 'w-12 h-12'
    }

    const thicknessVariants = {
        s: 'border-2',
        m: 'border-4',
        l: 'border-8'
    }

    const colorVariants = {
        primary: 'border-primary',
        secondary: 'border-secondary',
        grey: 'border-gray-300',
        typography: 'border-typography-main'
    }

    let sizeClass = sizeVariants[size] || sizeVariants['l']
    let bThicknessClass = thicknessVariants[thickness] || thicknessVariants['l']
    let bColorClass = colorVariants[color] || colorVariants['primary']

    let blockOrInline = $derived(isBlock ? "flex h-full w-full items-center justify-center" : "inline-block")
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

