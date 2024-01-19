<script lang="ts">
    import {onMount} from "svelte";

    export let navName:string = ""
    export let dest:string = ""
    export let menu:HTMLElement
    let element:HTMLElement

    onMount(() => {
        if(navName === "Vote"){
            element.addEventListener('click', () => {
                menu.classList.add("right-0")
                menu.classList.remove("-right-[75%]")
            })
        }
    })

    $: iconClass = ():string => {
        switch (navName){
            case "Chat":
                return "fa-message"
            case "Vote":
                return "fa-star voteNav"
            case "Results":
                return "fa-medal"
            case "Settings":
                return "fa-gear"
            default:
                return ""
        }
    }
</script>

<a href={dest} class="block">
    <div class="text-center">
        <span class="text-typography-nav" class:voteNav={navName === "Vote"} bind:this={element}>
            <span class="block pb-1">
                <i class="{iconClass()} fa-solid fa-2xl text-nav-icon"></i>
            </span>
            {navName}
        </span>
    </div>
</a>
