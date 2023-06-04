<script lang="ts">

    export let options:Array<{ label: string; value: number }>;
    export let legend:String;
    export let userSelected:number;
    export let name:String;


</script>

<h4>{legend}</h4>

{#each options as { value, label }}
    <input type="radio" bind:group={userSelected} value={value} name={name} on:change={() => this.parentElement.requestSubmit()}/>
    <label>{label}</label>
{/each}