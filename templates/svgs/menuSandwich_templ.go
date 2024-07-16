// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package svgs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func MenuButtonSVG(args templ.Attributes) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg id=\"svg2\" viewBox=\"0 0 300 300\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, args)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><defs id=\"defs4\"><linearGradient inkscape:collect=\"always\" id=\"linearGradient4168\"><stop style=\"stop-color:#53536c;stop-opacity:1\" offset=\"0\" id=\"stop4170\"></stop> <stop id=\"stop4176\" offset=\"0.6000002\" style=\"stop-color:#6f6f91;stop-opacity:1\"></stop> <stop style=\"stop-color:#b7b7c8;stop-opacity:1\" offset=\"1\" id=\"stop4172\"></stop></linearGradient> <linearGradient inkscape:collect=\"always\" id=\"linearGradient5871\"><stop style=\"stop-color:#f9f9f9;stop-opacity:1\" offset=\"0\" id=\"stop5873\"></stop> <stop style=\"stop-color:#e3e2db;stop-opacity:0\" offset=\"1\" id=\"stop5875\"></stop></linearGradient> <radialGradient r=\"16\" gradientTransform=\"matrix(1.5315,0,0,1.028,-4.3014,-3.2861)\" cx=\"15.286\" cy=\"33.369999\" gradientUnits=\"userSpaceOnUse\" id=\"radialGradient3622\"><stop offset=\"0\" stop-color=\"#fc0\" id=\"stop3624\" style=\"stop-color:#b7b7c8;stop-opacity:1\"></stop> <stop offset=\"1\" stop-color=\"#806600\" id=\"stop3626\" style=\"stop-color:#53536c;stop-opacity:1\"></stop></radialGradient> <radialGradient id=\"radialGradient3193\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,-8.1236,-2.5428)\" r=\"16\"><stop id=\"stop2845\" stop-color=\"#5aa02c\" offset=\"0\"></stop> <stop id=\"stop2847\" stop-color=\"#2d5016\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient3195\" x1=\"28.211\" xlink:href=\"#linearGradient2798\" gradientUnits=\"userSpaceOnUse\" y1=\"28.798\" gradientTransform=\"matrix(1.0345,0,0,1,-0.034483,0)\" x2=\"7.2297001\" y2=\"4.2326002\"></linearGradient> <linearGradient id=\"linearGradient2798\"><stop id=\"stop2800\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop2802\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient3197\" x1=\"20.540001\" xlink:href=\"#linearGradient2780\" gradientUnits=\"userSpaceOnUse\" y1=\"5.9130998\" gradientTransform=\"matrix(1.0334,0,0,0.9055,-0.58308,0.17787)\" x2=\"20.540001\" y2=\"2.5\"></linearGradient> <linearGradient id=\"linearGradient2780\"><stop id=\"stop2782\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop2784\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <radialGradient id=\"radialGradient3199\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,-4.3014,-3.2861)\" r=\"16\"><stop id=\"stop2889\" stop-color=\"#fc0\" offset=\"0\"></stop> <stop id=\"stop2891\" stop-color=\"#806600\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient3201\" x1=\"28.211\" xlink:href=\"#linearGradient2798\" gradientUnits=\"userSpaceOnUse\" y1=\"28.798\" gradientTransform=\"matrix(1.0345,0,0,1,3.7877,-0.74321)\" x2=\"7.2297001\" y2=\"4.2326002\"></linearGradient> <linearGradient id=\"linearGradient3214\"><stop id=\"stop3216\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop3218\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient3203\" x1=\"20.540001\" xlink:href=\"#linearGradient2780\" gradientUnits=\"userSpaceOnUse\" y1=\"5.9130998\" gradientTransform=\"matrix(1.0334,0,0,0.9055,3.2392,-0.56534)\" x2=\"20.540001\" y2=\"2.5\"></linearGradient> <linearGradient gradientUnits=\"userSpaceOnUse\" y2=\"87.845787\" x2=\"13.749455\" y1=\"87.845787\" x1=\"2.0156972\" gradientTransform=\"scale(2.556753,0.39112108)\" id=\"linearGradient3221\"><stop id=\"stop3223\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop3225\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <radialGradient id=\"radialGradient3205\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,27.975,-4.2416)\" r=\"16\"><stop id=\"stop2933\" stop-color=\"#5f8dd3\" offset=\"0\"></stop> <stop id=\"stop2935\" stop-color=\"#214478\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient3207\" x1=\"28.211\" xlink:href=\"#linearGradient2798\" gradientUnits=\"userSpaceOnUse\" y1=\"28.798\" gradientTransform=\"matrix(1.0345,0,0,1,36.064,-1.6988)\" x2=\"7.2297001\" y2=\"4.2326002\"></linearGradient> <linearGradient id=\"linearGradient3231\"><stop id=\"stop3233\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop3235\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient3209\" x1=\"20.540001\" xlink:href=\"#linearGradient2780\" gradientUnits=\"userSpaceOnUse\" y1=\"5.9130998\" gradientTransform=\"matrix(1.0334,0,0,0.9055,35.516,-1.5209)\" x2=\"20.540001\" y2=\"2.5\"></linearGradient> <linearGradient id=\"linearGradient3238\"><stop id=\"stop3240\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop3242\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <radialGradient id=\"radialGradient3211\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,-39.453,29.502)\" r=\"16\"><stop id=\"stop2971\" stop-color=\"#ff7f2a\" offset=\"0\"></stop> <stop id=\"stop2973\" stop-color=\"#803300\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient3213\" x1=\"28.211\" xlink:href=\"#linearGradient2798\" gradientUnits=\"userSpaceOnUse\" y1=\"28.798\" gradientTransform=\"matrix(1.0345,0,0,1,-31.364,32.045)\" x2=\"7.2297001\" y2=\"4.2326002\"></linearGradient> <linearGradient id=\"linearGradient3248\"><stop id=\"stop3250\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop3252\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient3215\" x1=\"20.540001\" xlink:href=\"#linearGradient2780\" gradientUnits=\"userSpaceOnUse\" y1=\"5.9130998\" gradientTransform=\"matrix(1.0334,0,0,0.9055,-31.912,32.223)\" x2=\"20.540001\" y2=\"2.5\"></linearGradient> <linearGradient id=\"linearGradient3255\"><stop id=\"stop3257\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop3259\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <radialGradient id=\"radialGradient3217\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,-3.9699,29.405)\" r=\"16\"><stop id=\"stop3009\" stop-color=\"#f00\" offset=\"0\"></stop> <stop id=\"stop3011\" stop-color=\"#800000\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient3219\" x1=\"28.211\" xlink:href=\"#linearGradient2798\" gradientUnits=\"userSpaceOnUse\" y1=\"28.798\" gradientTransform=\"matrix(1.0345,0,0,1,4.1192,31.948)\" x2=\"7.2297001\" y2=\"4.2326002\"></linearGradient> <linearGradient id=\"linearGradient3265\"><stop id=\"stop3267\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop3269\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient3271\" x1=\"20.540001\" xlink:href=\"#linearGradient2780\" gradientUnits=\"userSpaceOnUse\" y1=\"5.9130998\" gradientTransform=\"matrix(1.0334,0,0,0.9055,3.5706,32.126)\" x2=\"20.540001\" y2=\"2.5\"></linearGradient> <linearGradient id=\"linearGradient3273\"><stop id=\"stop3275\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop3277\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <radialGradient id=\"radialGradient3223\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,30.875,28.565)\" r=\"16\"><stop id=\"stop3047\" stop-color=\"#666\" offset=\"0\"></stop> <stop id=\"stop3049\" stop-color=\"#333\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient3225\" x1=\"28.211\" xlink:href=\"#linearGradient2798\" gradientUnits=\"userSpaceOnUse\" y1=\"28.798\" gradientTransform=\"matrix(1.0345,0,0,1,38.964,31.107)\" x2=\"7.2297001\" y2=\"4.2326002\"></linearGradient> <linearGradient id=\"linearGradient3283\"><stop id=\"stop3285\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop3287\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient3227\" x1=\"20.540001\" xlink:href=\"#linearGradient2780\" gradientUnits=\"userSpaceOnUse\" y1=\"5.9130998\" gradientTransform=\"matrix(1.0334,0,0,0.9055,38.415,31.285)\" x2=\"20.540001\" y2=\"2.5\"></linearGradient> <linearGradient id=\"linearGradient3290\"><stop id=\"stop3292\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop3294\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <radialGradient id=\"radialGradient3199-3\" gradientUnits=\"userSpaceOnUse\" cy=\"33.369999\" cx=\"15.286\" gradientTransform=\"matrix(1.5315,0,0,1.028,-4.3014,-3.2861)\" r=\"16\"><stop id=\"stop2889-5\" stop-color=\"#fc0\" offset=\"0\"></stop> <stop id=\"stop2891-6\" stop-color=\"#806600\" offset=\"1\"></stop></radialGradient> <linearGradient id=\"linearGradient2798-2\"><stop id=\"stop2800-9\" stop-color=\"#fff\" stop-opacity=\".19417\" offset=\"0\"></stop> <stop id=\"stop2802-1\" stop-color=\"#fff\" stop-opacity=\"0\" offset=\"1\"></stop></linearGradient> <linearGradient id=\"linearGradient2780-2\"><stop id=\"stop2782-7\" stop-opacity=\"0\" offset=\"0\"></stop> <stop id=\"stop2784-0\" stop-color=\"#fff\" stop-opacity=\".23301\" offset=\"1\"></stop></linearGradient> <filter inkscape:label=\"Drop Shadow\" id=\"filter5201\" color-interpolation-filters=\"sRGB\"><feFlood flood-opacity=\"0.5\" flood-color=\"rgb(0,0,0)\" result=\"flood\" id=\"feFlood5203\"></feFlood> <feComposite in=\"flood\" in2=\"SourceGraphic\" operator=\"in\" result=\"composite1\" id=\"feComposite5205\"></feComposite> <feGaussianBlur stdDeviation=\"6\" result=\"blur\" id=\"feGaussianBlur5207\"></feGaussianBlur> <feOffset dx=\"21\" dy=\"5\" result=\"offset\" id=\"feOffset5209\"></feOffset> <feComposite in=\"SourceGraphic\" in2=\"offset\" operator=\"over\" result=\"composite2\" id=\"feComposite5211\"></feComposite></filter> <filter color-interpolation-filters=\"sRGB\" id=\"C\"><feGaussianBlur stdDeviation=\"2.58594\" id=\"r\"></feGaussianBlur></filter> <radialGradient xlink:href=\"#8\" r=\"44.419998\" cy=\"100.32\" cx=\"81.790001\" gradientTransform=\"matrix(2.0026052,-1.5970314,1.7773333,2.2286954,-278.08797,27.013826)\" gradientUnits=\"userSpaceOnUse\" id=\"F\"></radialGradient> <linearGradient id=\"8\"><stop id=\"k\" stop-color=\"#eee\"></stop> <stop offset=\"1\" id=\"l\" stop-color=\"#d2d2d2\"></stop></linearGradient> <radialGradient xlink:href=\"#9\" id=\"D\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(1,0,0,0.97467889,0,0.43910718)\" cx=\"89.510002\" cy=\"22.254\" r=\"18.278999\"></radialGradient> <linearGradient id=\"9\"><stop id=\"m\"></stop> <stop offset=\"1\" id=\"n\" stop-opacity=\"0.536\"></stop></linearGradient> <filter color-interpolation-filters=\"sRGB\" id=\"B\" x=\"-0.147\" width=\"1.294\" y=\"-0.145\" height=\"1.29\"><feGaussianBlur stdDeviation=\"1.81881\" id=\"q\"></feGaussianBlur></filter> <linearGradient xlink:href=\"#A\" y2=\"25.1\" x2=\"86.82\" y1=\"14.793\" x1=\"98.330002\" gradientTransform=\"matrix(0.96714879,0,0,0.96714879,1.5886796,3.2231964)\" gradientUnits=\"userSpaceOnUse\" id=\"E\"></linearGradient> <linearGradient id=\"A\"><stop id=\"o\"></stop> <stop offset=\"1\" id=\"p\" stop-color=\"#ddd\" stop-opacity=\"0\"></stop></linearGradient> <filter inkscape:label=\"Drop Shadow\" id=\"filter5664\" color-interpolation-filters=\"sRGB\"><feFlood flood-opacity=\"0.6\" flood-color=\"rgb(0,0,0)\" result=\"flood\" id=\"feFlood5666\"></feFlood> <feComposite in=\"flood\" in2=\"SourceGraphic\" operator=\"in\" result=\"composite1\" id=\"feComposite5668\"></feComposite> <feGaussianBlur stdDeviation=\"6\" result=\"blur\" id=\"feGaussianBlur5670\"></feGaussianBlur> <feOffset dx=\"6\" dy=\"6\" result=\"offset\" id=\"feOffset5672\"></feOffset> <feComposite in=\"SourceGraphic\" in2=\"offset\" operator=\"over\" result=\"composite2\" id=\"feComposite5674\"></feComposite></filter> <filter inkscape:label=\"Drop Shadow\" id=\"filter5827\" color-interpolation-filters=\"sRGB\"><feFlood flood-opacity=\"0.6\" flood-color=\"rgb(0,0,0)\" result=\"flood\" id=\"feFlood5829\"></feFlood> <feComposite in=\"flood\" in2=\"SourceGraphic\" operator=\"in\" result=\"composite1\" id=\"feComposite5831\"></feComposite> <feGaussianBlur stdDeviation=\"1\" result=\"blur\" id=\"feGaussianBlur5833\"></feGaussianBlur> <feOffset dx=\"1\" dy=\"1\" result=\"offset\" id=\"feOffset5835\"></feOffset> <feComposite in=\"SourceGraphic\" in2=\"offset\" operator=\"over\" result=\"composite2\" id=\"feComposite5837\"></feComposite></filter> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient5871\" id=\"radialGradient5877\" cx=\"-1918.6243\" cy=\"-1878.2841\" fx=\"-1918.6243\" fy=\"-1878.2841\" r=\"53.47414\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(0.56626527,-1.9491737,2.6535343,0.77089145,5592.9315,-3935.7797)\"></radialGradient> <filter inkscape:label=\"Drop Shadow\" id=\"filter6216\" color-interpolation-filters=\"sRGB\"><feFlood flood-opacity=\"0.5\" flood-color=\"rgb(0,0,0)\" result=\"flood\" id=\"feFlood6218\"></feFlood> <feComposite in=\"flood\" in2=\"SourceGraphic\" operator=\"in\" result=\"composite1\" id=\"feComposite6220\"></feComposite> <feGaussianBlur stdDeviation=\"10\" result=\"blur\" id=\"feGaussianBlur6222\"></feGaussianBlur> <feOffset dx=\"-10\" dy=\"10\" result=\"offset\" id=\"feOffset6224\"></feOffset> <feComposite in=\"SourceGraphic\" in2=\"offset\" operator=\"over\" result=\"composite2\" id=\"feComposite6226\"></feComposite></filter> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient5871\" id=\"radialGradient6309\" cx=\"-669.98749\" cy=\"-2476.3098\" fx=\"-669.98749\" fy=\"-2476.3098\" r=\"121.94899\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(0.63293951,-0.49176645,0.39982869,0.51460622,1015.3149,-871.82865)\"></radialGradient> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient5871\" id=\"radialGradient3127\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(0.63293951,-0.49176645,0.39982869,0.51460622,1015.3149,-871.82865)\" cx=\"-669.98749\" cy=\"-2476.3098\" fx=\"-669.98749\" fy=\"-2476.3098\" r=\"121.94899\"></radialGradient> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient5871\" id=\"radialGradient3129\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(0.56626527,-1.9491737,2.6535343,0.77089145,5592.9315,-3935.7797)\" cx=\"-1918.6243\" cy=\"-1878.2841\" fx=\"-1918.6243\" fy=\"-1878.2841\" r=\"53.47414\"></radialGradient> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient5871\" id=\"radialGradient3128\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(0.56626527,-1.9491737,2.6535343,0.77089145,5592.9315,-3935.7797)\" cx=\"-1918.6243\" cy=\"-1878.2841\" fx=\"-1918.6243\" fy=\"-1878.2841\" r=\"53.47414\"></radialGradient> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient5871\" id=\"radialGradient3136\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"matrix(0.67830627,-0.52701445,0.42848693,0.55149129,1656.5395,2090.3276)\" cx=\"-669.98749\" cy=\"-2476.3098\" fx=\"-669.98749\" fy=\"-2476.3098\" r=\"121.94899\"></radialGradient> <radialGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4204\" id=\"radialGradient4210\" cx=\"350.02145\" cy=\"458.27914\" fx=\"350.02145\" fy=\"458.27914\" r=\"54.11713\" gradientTransform=\"matrix(1.2234571e-7,-4.010857,1.1502599,3.5087105e-8,-177.11872,1862.163)\" gradientUnits=\"userSpaceOnUse\"></radialGradient> <linearGradient inkscape:collect=\"always\" id=\"linearGradient4204\"><stop style=\"stop-color:#1a1a1a;stop-opacity:1;\" offset=\"0\" id=\"stop4206\"></stop> <stop style=\"stop-color:#1a1a1a;stop-opacity:0;\" offset=\"1\" id=\"stop4208\"></stop></linearGradient> <linearGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4168\" id=\"linearGradient4174\" x1=\"146.26012\" y1=\"864.08765\" x2=\"146.26012\" y2=\"830.97693\" gradientUnits=\"userSpaceOnUse\"></linearGradient> <linearGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4168\" id=\"linearGradient4180\" gradientUnits=\"userSpaceOnUse\" x1=\"146.26012\" y1=\"864.08765\" x2=\"146.26012\" y2=\"830.97693\" gradientTransform=\"translate(0,51.287415)\"></linearGradient> <linearGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4168\" id=\"linearGradient4184\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"translate(0,102.57642)\" x1=\"146.26012\" y1=\"864.08765\" x2=\"146.26012\" y2=\"830.97693\"></linearGradient> <linearGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4168\" id=\"linearGradient4191\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"translate(0,51.287415)\" x1=\"146.26012\" y1=\"864.08765\" x2=\"146.26012\" y2=\"830.97693\"></linearGradient> <linearGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4168\" id=\"linearGradient4193\" gradientUnits=\"userSpaceOnUse\" gradientTransform=\"translate(0,102.57642)\" x1=\"146.26012\" y1=\"864.08765\" x2=\"146.26012\" y2=\"830.97693\"></linearGradient> <linearGradient inkscape:collect=\"always\" xlink:href=\"#linearGradient4168\" id=\"linearGradient4195\" gradientUnits=\"userSpaceOnUse\" x1=\"146.26012\" y1=\"864.08765\" x2=\"146.26012\" y2=\"830.97693\"></linearGradient></defs> <sodipodi:namedview id=\"base\" pagecolor=\"#ffffff\" bordercolor=\"#666666\" borderopacity=\"1.0\" inkscape:pageopacity=\"0.0\" inkscape:pageshadow=\"2\" inkscape:zoom=\"1.7962932\" inkscape:cx=\"149.99999\" inkscape:cy=\"149.99992\" inkscape:document-units=\"px\" inkscape:current-layer=\"layer1\" showgrid=\"false\" showguides=\"true\" inkscape:guide-bbox=\"true\" inkscape:window-width=\"1920\" inkscape:window-height=\"1025\" inkscape:window-x=\"0\" inkscape:window-y=\"33\" inkscape:window-maximized=\"1\"></sodipodi:namedview> <g inkscape:label=\"Camada 1\" inkscape:groupmode=\"layer\" id=\"layer1\" transform=\"translate(0,-752.36218)\"><g id=\"g4197\"><rect ry=\"30.000004\" rx=\"29.999998\" y=\"771.67261\" x=\"19.310228\" height=\"261.3797\" width=\"261.3797\" id=\"rect6270\" style=\"fill:url(#radialGradient3136);fill-opacity:1;stroke:none\"></rect> <rect transform=\"matrix(3.173314,0,0,3.173314,1777.8554,6245.1647)\" style=\"opacity:0.6;fill:url(#radialGradient3128);fill-opacity:1;stroke:#cccccc;stroke-width:0.63025582;stroke-miterlimit:4;stroke-dasharray:none;filter:url(#filter5827)\" id=\"rect5797\" width=\"88.123566\" height=\"88.123566\" x=\"-557.04456\" y=\"-1727.7283\" rx=\"9.4538393\" ry=\"9.4538393\"></rect> <g transform=\"translate(-3.1994699e-4,-9.4034393e-5)\" id=\"g4186\"><rect rx=\"13.430069\" ry=\"13.430069\" style=\"fill:url(#linearGradient4191);fill-opacity:1;stroke:none\" id=\"rect4178\" width=\"161.16084\" height=\"33.575172\" x=\"69.419884\" y=\"885.57397\"></rect> <rect y=\"936.86298\" x=\"69.419884\" height=\"33.575172\" width=\"161.16084\" id=\"rect4182\" style=\"fill:url(#linearGradient4193);fill-opacity:1;stroke:none\" ry=\"13.430069\" rx=\"13.430069\"></rect> <rect y=\"834.28656\" x=\"69.419884\" height=\"33.575172\" width=\"161.16084\" id=\"rect3659\" style=\"fill:url(#linearGradient4195);fill-opacity:1;stroke:none\" ry=\"13.430069\" rx=\"13.430069\"></rect></g></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}