<!-- @migration-task Error while migrating Svelte code: can't migrate `let isDisabled:boolean = true` to `$state` because there's a variable named state.
     Rename the variable and try again or migrate by hand. -->
<script lang="ts">
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import Spinner from "$lib/components/Spinner.svelte";

    export let buttonState:number
    let isDisabled:boolean = true

    $: if (buttonState){
        switch (buttonState) {
            case formButtonState.ENABLED:
                isDisabled = false
                break
            default:
                isDisabled = true
        }
    }

    $: disabledClass = isDisabled ? "bg-gray-500 cursor-not-allowed" : ""
</script>


<button type="submit" disabled={isDisabled} class="{disabledClass} flex" >
    {#if buttonState === formButtonState.SENDING}
        <Spinner size={"sm"} thickness={"s"} />
    {:else}
        <slot/>
    {/if}
</button>