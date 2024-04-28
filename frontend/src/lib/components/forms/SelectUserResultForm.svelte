<script lang="ts">
    import {resultPageState} from "$lib/stores/resultPageState.store.js";
    import {authLvl} from "$lib/models/enums/authLvl.enum.js";
    import { enhance } from '$app/forms';
    import {userStore} from "$lib/stores/user.store";
    import {voteCats} from "$lib/models/enums/categories.enum";

    export let showTitle = false
    let form:HTMLFormElement
    let userArray = [...new Map(Object.entries($userStore))]
    let currentCat

    const changeCat = () => {
        $resultPageState.category = currentCat
    }
</script>

{#if showTitle}
    <h3 class="text-center">Filters</h3>
{/if}

<div class="pb-3">
    <form method="POST" action="?/getUserResults" use:enhance bind:this={form}>
        <div class="flex">
            <div class="w-1/4 text-right pt-2 pr-3">
                <label for="id">Person</label>
            </div>
            <div class="w-1/2">
                <select class="w-full text-center py-2"
                        name="id"
                        bind:value={$resultPageState.userId}
                        on:change={() => form.requestSubmit()}>
                    <option value="">Everyone</option>
                    {#each userArray as userInfo}
                        {#if userInfo[1].authLvl !== authLvl.BOT}
                            <option value={userInfo[0]}>{userInfo[1].name}</option>
                        {/if}
                    {/each}
                </select>
            </div>
        </div>
    </form>
</div>

<div class="pb-3">
    <div class="flex">
        <div class="w-1/4 text-right pt-2 pr-3">
            <label for="id">Category</label>
        </div>
        <div class="w-1/2">
            <select class="w-full text-center py-2 capitalize"
                    name="id"
                    bind:value={currentCat}
                    on:change={changeCat}>
                <option selected={$resultPageState.category === voteCats.TOTAL}
                        value={voteCats.TOTAL}>
                    Total
                </option>

                <option selected={$resultPageState.category === voteCats.SONG}
                        value={voteCats.SONG}>
                    Song
                </option>
                <option selected={$resultPageState.category === voteCats.PERFORMANCE}
                        value={voteCats.PERFORMANCE}>
                    Performance
                </option>
                <option selected={$resultPageState.category === voteCats.COSTUME}
                        value={voteCats.COSTUME}>
                    Costume
                </option>
                <option selected={$resultPageState.category === voteCats.PROPS}
                        value={voteCats.PROPS}>
                    Props
                </option>
            </select>
        </div>
    </div>
</div>