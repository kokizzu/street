<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { FaSolidArrowRight } from './node_modules/svelte-icons-pack/dist/fa';
  
  let error = '#{error}';
  function backHome() {
    window.location.href = "/"
  }
</script>

<div class="error_container">
	<header>
		<h1>Error</h1>
		<h2>{error}</h2>
	</header>
	<button class="back_home" on:click={backHome}>
		<span>Back to home</span>
		<Icon className="arrow_backhome" color='#FFF' size={18} src={FaSolidArrowRight}/>
	</button>
</div>

<style>
	.error_container {
		 display: flex;
		 flex-direction: column;
		 justify-content: center;
		 align-items: center;
		 gap: 40px;
		 width: fit-content;
		 height: fit-content;
		 margin: 100px auto auto auto;
		 color: #475569;
	}
   .error_container header {
	    display: flex;
	    flex-direction: column;
	    gap: 15px;
	    text-align: center;
   }
   .error_container header h1 {
	    margin: 0;
	    font-size: 80px;
   }
   .error_container header h2 {
	    margin: 0;
	    font-size: 25px;
	    font-weight: 500;
   }
   .error_container .back_home {
	    position: relative;
	    width: fit-content;
	    padding: 20px 50px 20px 30px;
	    background-color: #3b82f6;
	    border: none;
	    border-radius: 9999px;
	    font-size: 18px;
	    cursor: pointer;
   }
   .error_container .back_home:hover {
	    background-color: #2C74EA;
   }
   .error_container .back_home span {
	    color: #FFFFFF;
   }
   :global(.arrow_backhome) {
	    position: absolute;
	    right: 22px;
	    top: auto;
	    bottom: auto;
	    margin-top : 2px;
       transition-property: all;
       transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
       transition-duration: 150ms;
   }
   :global(.error_container .back_home:hover .arrow_backhome) {
	    right: 16px;
   }
</style>