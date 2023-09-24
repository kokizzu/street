<script>
  import Footer from './_components/Footer.svelte';
</script>

<section class="tos_container">
	<header>
		<h1>HapSTR Website Terms and Conditions of Use</h1>
	</header>
	<div>
		<ol>
			<li>
				<span>Terms</span>
				<p>
					By accessing this Website, accessible from hapstr.xyz, you are agreeing to be bound by these Website Terms and
					Conditions of Use and agree that you are responsible for the agreement with any applicable local laws. If you
					disagree with any of these terms, you are prohibited from accessing this site. The materials contained in this
					Website are protected by copyright and trade mark law.
				</p>
			</li>
			<li>
				<span>Use License</span>
				<p>
					Permission is granted to temporarily download one copy of the materials on HapSTR's Website for personal,
					non-commercial transitory viewing only. This is the grant of a license, not a transfer of title, and under
					this license you may not:
				</p>
				<ul>
					<li>Modify or copy the materials;</li>
					<li>Use the materials for any commercial purpose or for any public display;</li>
					<li>Attempt to reverse engineer any software contained on HapSTR's Website;</li>
					<li>Remove any copyright or other proprietary notations from the materials;</li>
					<li>Transferring the materials to another person or "mirror" the materials on any other server.</li>
				</ul>
				<p>
					This will let HapSTR to terminate upon violations of any of these restrictions. Upon termination, your viewing
					right will also be terminated and you should destroy any downloaded materials in your possession whether it is
					printed or electronic format.
				</p>
			</li>
			<li>
				<span>Disclaimer</span>
				<p>
					All the materials on HapSTR's Website are provided “as is”. HapSTR makes no warranties, may it be expressed or
					implied, therefore negates all other warranties. Furthermore, HapSTR does not make any representations
					concerning the accuracy or reliability of the use of the materials on its Website or otherwise relating to
					such materials or any sites linked to this Website.
				</p>
			</li>
			<li>
				<span>Limitations</span>
				<p>
					HapSTR or its suppliers will not be hold accountable for any damages that will arise with the use or inability
					to use the materials on HapSTR's Website, even if HapSTR or an authorize representative of this Website has
					been notified, orally or written, of the possibility of such damage. Some jurisdiction does not allow
					limitations on implied warranties or limitations of liability for incidental damages, these limitations may
					not apply to you.
				</p>
			</li>
			<li>
				<span>Revisions and Errata</span>
				<p>
					The materials appearing on HapSTR's Website may include technical, typographical, or photographic errors.
					HapSTR will not promise that any of the materials in this Website are accurate, complete, or current. HapSTR
					may change the materials contained on its Website at any time without notice. HapSTR does not make any
					commitment to update the materials.
				</p>
			</li>
			<li>
				<span>Links</span>
				<p>
					HapSTR has not reviewed all of the sites linked to its Website and is not responsible for the contents of any
					such linked site. The presence of any link does not imply endorsement by HapSTR of the site. The use of any
					linked website is at the user's own risk.
				</p>
			</li>
			<li>
				<span>Site Terms of Use Modifications</span>
				<p>
					HapSTR may revise these Terms of Use for its Website at any time without prior notice. By using this Website,
					you are agreeing to be bound by the current version of these Terms and Conditions of Use.
				</p>
			</li>
			<li>
				<span>Governing Law</span>
				<p>
					Any claim related to HapSTR's Website shall be governed by the laws of Country without regards to its conflict
					of law provisions.
				</p>
			</li>
		</ol>
	</div>
	<Footer/>
</section>

<style>
    .tos_container {
        display        : flex;
        flex-direction : column;
        width          : 780px;
        margin         : 0 auto;
        color          : #475569;
    }

    .tos_container header {
        margin : 30px auto 20px auto;
    }

    .tos_container header h1 {
        font-size : 30px;
        margin    : 0;
    }

    .tos_container div ol li::marker {
        font-size   : 18px;
        font-weight : bold;
    }

    .tos_container div ol li span {
        font-size   : 20px;
        font-weight : bold;
    }
</style>
