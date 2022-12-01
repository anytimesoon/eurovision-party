<script type="ts">
	import type { Load } from '@sveltejs/kit';
	import type { TokenModel } from '$lib/models/classes/token.model';
	import { setToken, tokenStore } from '$lib/stores/token.store';
	import {onMount} from "svelte";

	onMount( async () => {
		let path = window.location.pathname;
		let params = path.split('/');

		fetch("http://localhost:8080/login", {
			method: "POST",
			body: JSON.stringify({
				userId: params[3],
				token: params[2]
			}),
		}).
		then(res => res.json() as Promise<TokenModel>).
		then(token => setToken(token)).
		catch((e) => {
			console.log(e)
		})
	});

	let token:TokenModel
	tokenStore.subscribe( (val) => {
		token = val
	})	
</script>

<div>
	{#if token != null}
		{token.token}
	{/if}
</div>