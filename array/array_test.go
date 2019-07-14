package array

import (
	"testing"
	"fmt"
)

// go test -v .\array -test.run TestMaxSubArray
func TestMaxSubArray(t *testing.T) {
	nums := []int{2, 1}
	v := maxSubArray(nums)
	fmt.Println(v)
}

// go test -v .\array -test.run TestGenerate
func TestGenerate(t *testing.T) {
	v := generate(5)
	fmt.Println(v)
}

func TestGetRow(t *testing.T) {
	v := getRow(3)
	fmt.Println(v)
}

func TestMaxProfit(t *testing.T) {
	v := maxProfit([]int{7, 1, 5, 3, 6, 4})

	fmt.Println(v)
}

// [1,2,3,4,4,9,56,90]

func TestTwoSum(t *testing.T) {
	v := twoSum([]int{2, 7, 11, 15}, 18)

	fmt.Println(v)
}

// [1,2,3,4,4,9,56,90]

func TestRotate(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(v, 3)
	fmt.Println(v)
}

func TestContainsNearbyDuplicate(t *testing.T) {
	v := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	fmt.Println(v)
}

func TestFindDisappearedNumbers(t *testing.T) {
	v := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(v)
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	v := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})

	fmt.Println(v)
}

func TestFib(t *testing.T) {
	v := fib(4)
	fmt.Println(v)
}

func TestFindPairs(t *testing.T) {
	// v := findPairs([]int{3, 1, 4, 1, 5}, 2)
	v := findPairs([]int{0, 0, 1, 0, 0}, 1)
	fmt.Println(v)
}

func TestArrayPairSum(t *testing.T) {
	// v := findPairs([]int{1,4,3,2})
	v := arrayPairSum([]int{-1, 4, 3, -2})
	fmt.Println(v)
}

func TestFindUnsortedSubarray(t *testing.T) {
	v := findUnsortedSubarray([]int{1, 3, 2, 2, 2})
	fmt.Println(v)
}

func TestMaximumProduct(t *testing.T) {
	v := maximumProduct([]int{722, 634, -504, -379, 163, -613, -842, -578, 750, 951, -158, 30, -238, -392, -487, -797, -157, -374, 999, -5, -521, -879, -858, 382, 626, 803, -347, 903, -205, 57, -342, 186, -736, 17, 83, 726, -960, 343, -984, 937, -758, -122, 577, -595, -544, -559, 903, -183, 192, 825, 368, -674, 57, -959, 884, 29, -681, -339, 582, 969, -95, -455, -275, 205, -548, 79, 258, 35, 233, 203, 20, -936, 878, -868, -458, -882, 867, -664, -892, -687, 322, 844, -745, 447, -909, -586, 69, -88, 88, 445, -553, -666, 130, -640, -918, -7, -420, -368, 250, -786})
	fmt.Println(v)
	// 943695360
}

func TestFindMaxAverage(t *testing.T) {
	v := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	fmt.Println(v)
}

func TestFindShortestSubArray(t *testing.T) {
	v := findShortestSubArray([]int{1, 2})
	fmt.Println(v)
}

func TestA(t *testing.T) {
	v := sortArrayByParity([]int{4279, 4333, 4035, 914, 3650, 590, 1267, 1101, 4412, 4854, 4766, 2809, 65, 199, 3813, 4696, 3131, 2873, 1210, 1191, 2417, 3036, 708, 1405, 4340, 1469, 4693, 1972, 805, 3088, 1734, 3112, 274, 1269, 3566, 971, 2332, 2748, 4854, 467, 3395, 3754, 2107, 3385, 4081, 1737, 3386, 3533, 4260, 49, 536, 2934, 4078, 3552, 2592, 2344, 2011, 2325, 805, 1625, 3517, 1608, 4326, 466, 2348, 2446, 2503, 3542, 648, 1490, 889, 4963, 4902, 927, 3882, 1123, 564, 1321, 4317, 3134, 1429, 604, 329, 2828, 4609, 1267, 2589, 2931, 1357, 346, 3454, 2087, 285, 3647, 3367, 1061, 1956, 4265, 3869, 3653, 332, 182, 4744, 3012, 2527, 1968, 1758, 70, 962, 4203, 4663, 3062, 3692, 916, 170, 2936, 1051, 499, 3424, 3050, 1175, 4590, 3645, 2792, 2817, 4399, 4658, 1359, 1796, 1826, 152, 3346, 4335, 3666, 3743, 22, 118, 856, 1503, 3764, 1047, 3110, 1518, 985, 3885, 4267, 310, 729, 711, 584, 4999, 1764, 1287, 3539, 1529, 2656, 1041, 3718, 4026, 2693, 1243, 2273, 2660, 4278, 693, 2871, 1239, 4829, 4279, 4162, 3563, 3605, 2671, 4629, 82, 2789, 510, 4895, 2495, 3011, 939, 4898, 1828, 1341, 2423, 1675, 2417, 161, 3646, 3765, 4167, 4462, 1964, 4342, 1609, 1639, 1325, 4967, 492, 4021, 1888, 2037, 3544, 3732, 4318, 3118, 1781, 2732, 995, 141, 1438, 1074, 2581, 999, 1675, 3309, 31, 3208, 3205, 4216, 1857, 564, 4976, 1962, 4519, 4971, 2371, 4528, 862, 691, 2773, 3429, 261, 1852, 1474, 551, 2840, 4967, 4562, 2605, 2185, 2011, 3090, 457, 3344, 4029, 564, 1820, 3711, 1053, 2033, 2016, 3413, 4163, 1463, 3333, 4483, 229, 2016, 284, 1890, 3570, 1919, 1131, 471, 4989, 4602, 3517, 2429, 4223, 4391, 2955, 2508, 204, 3375, 209, 1550, 3981, 4367, 4781, 1164, 990, 432, 3663, 4668, 4734, 1367, 3337, 3687, 3911, 4731, 852, 1937, 363, 3103, 1926, 229, 1054, 915, 4491, 2011, 4194, 361, 4112, 3390, 88, 3424, 1998, 4813, 1255, 2184, 2410, 41, 4796, 344, 4421, 4495, 4802, 1817, 401, 710, 458, 4401, 1729, 4049, 3800, 3084, 112, 221, 1264, 494, 3707, 2811, 4534, 4079, 1033, 4569, 2527, 3189, 1922, 2175, 2730, 95, 4461, 696, 784, 4552, 680, 524, 955, 4690, 809, 1075, 1715, 2058, 3418, 2574, 4025, 784, 4693, 2809, 1556, 3850, 3615, 1886, 4536, 1369, 4387, 57, 4130, 1385, 1743, 2349, 2070, 780, 2725, 1210, 2443, 4644, 3063, 916, 699, 3254, 234, 4513, 1171, 4937, 4291, 202, 3717, 1893, 3144, 2595, 4424, 3264, 795, 940, 1748, 315, 1312, 4758, 2801, 920, 1625, 2661, 244, 4932, 2973, 4275, 1618, 277, 3277, 1175, 3911, 4248, 1826, 18, 988, 1753, 710, 1798, 4000, 1770, 768, 1657, 3732, 3252, 3240, 4252, 3670, 2621, 4675, 4769, 3461, 3393, 2626, 1153, 58, 4663, 4397, 3900, 4249, 3464, 1114, 2219, 13, 1736, 3474, 1748, 4874, 4691, 2339, 2810, 55, 2763, 3119, 4314, 2775, 4681, 358, 4960, 968, 3545, 3996, 4670, 1549, 4302, 4127, 4573, 1804, 1726, 1948, 682, 1707, 421, 1545, 956, 4466, 1154, 1468, 3258, 239, 4854, 1134, 4694, 7, 2335, 2813, 1119, 3433, 351, 4525, 100, 3763, 3441, 974, 2577, 3850, 982, 635, 3403, 1045, 2810, 2867, 4571, 174, 732, 2672, 4200, 1462, 3655, 2629, 2424, 1401, 1709, 4635, 184, 2664, 4025, 973, 4904, 4173, 4749, 2960, 4273, 688, 166, 714, 304, 4494, 2845, 1704, 3983, 1173, 23, 1060, 4166, 1601, 2341, 1535, 1479, 3703, 814, 4601, 2902, 2695, 4042, 4685, 989, 1951, 2959, 3498, 4484, 4073, 729, 4820, 3474, 640, 1004, 2784, 3222, 4976, 3494, 3890, 1836, 4146, 115, 2209, 4691, 3811, 433, 3777, 2698, 1144, 3114, 4753, 1442, 450, 488, 1249, 1740, 427, 1997, 4917, 2337, 2347, 253, 758, 4953, 539, 3745, 2969, 1149, 2533, 4136, 3305, 337, 1029, 276, 1411, 3012, 2006, 2801, 544, 2478, 4947, 982, 2114, 4742, 269, 1404, 4594, 3444, 4447, 3901, 934, 2043, 65, 3419, 4701, 3575, 406, 3836, 4796, 1949, 1351, 4523, 1776, 4679, 4030, 1767, 1759, 4929, 2828, 483, 4285, 4426, 1181, 35, 3031, 4321, 2412, 4483, 2127, 4569, 4164, 1960, 4170, 4099, 217, 1798, 1273, 996, 3943, 1362, 4747, 4300, 681, 2563, 4775, 4876, 1534, 3462, 1846, 270, 3409, 315, 640, 3301, 3916, 780, 3232, 683, 1200, 1556, 188, 3716, 454, 2378, 4532, 3293, 4588, 892, 2241, 268, 3771, 3973, 3191, 1045, 2785, 2413, 2789, 4405, 4262, 4855, 2060, 3411, 2849, 238, 3881, 1605, 2698, 2069, 2896, 1422, 4016, 3980, 2445, 1533, 835, 89, 4185, 297, 164, 4391, 234, 2234, 2270, 665, 3602, 674, 3755, 2185, 2339, 2464, 872, 267, 4412, 2647, 4895, 3706, 335, 1972, 3281, 4074, 3499, 4997, 4301, 4875, 678, 1686, 3221, 4497, 3550, 4977, 2061, 4777, 665, 3563, 759, 1231, 3231, 3811, 565, 4189, 2892, 4060, 3276, 4832, 4695, 715, 2100, 2133, 3264, 1718, 1925, 3535, 3534, 4861, 3117, 1872, 3078, 4924, 1268, 1328, 2505, 780, 3093, 2269, 2928, 1778, 2007, 1740, 1097, 4862, 211, 1121, 4408, 690, 1950, 4819, 1213, 292, 4879, 2700, 999, 3256, 4524, 4678, 2118, 1722, 4732, 1558, 1593, 466, 429, 4876, 2614, 4666, 1755, 323, 2587, 1360, 950, 3184, 4336, 2017, 945, 4612, 1770, 2766, 203, 4540, 3621, 438, 4503, 1536, 936, 2201, 331, 4369, 4215, 3522, 1049, 4250, 3701, 2077, 865, 4281, 107, 3174, 4265, 4876, 3810, 2789, 3616, 1827, 2643, 2841, 4080, 3867, 2272, 3946, 1601, 636, 2788, 4314, 440, 377, 3206, 4869, 4719, 3581, 2155, 3518, 4424, 1282, 2961, 1005, 2174, 2573, 945, 324, 252, 776, 3531, 1210, 379, 724, 2585, 3102, 4684, 1671, 3208, 1874, 4774, 4501, 1731, 2094, 2680, 494, 4794, 1318, 3864, 2038, 4576, 2072, 343, 826, 273, 461, 1613, 754, 874, 242, 500, 2935, 1369, 2421, 1971, 4088, 1147, 419, 4250, 4505, 3414, 1056, 4986, 1191, 4413, 4056, 3645, 583, 2408, 770, 4403, 2119, 1649, 436, 2042, 497, 2307, 627, 3732, 585, 818, 628, 1398, 1434, 1121, 1324, 4311, 3292, 1805, 1406, 2913, 749, 772, 4281, 4492, 1378, 1937, 4357, 2298, 3215, 2185, 3461, 459, 2068, 1096, 1108, 2482, 2348, 451, 4314, 1295, 3918, 2065, 2735, 3994, 1095, 3001, 3123, 4417, 4122, 4731, 943, 3834, 2218, 1185, 804, 4260, 1049, 371, 4898, 1043, 2807, 3091, 2108, 930, 565, 481, 375, 1131, 621, 3097, 1091, 973, 4308, 2805, 3148, 4458, 1583, 2313, 1225, 1295, 2569, 1942, 725, 2313, 1682, 1985, 4947, 4784, 4519, 2450, 1021, 1280, 3665, 4882, 4683, 2501, 1404, 4786, 4220, 1132, 2360, 1104, 3720, 3020, 2910, 118, 1587, 1334, 1590, 3105, 523, 622, 3373, 1755, 3971, 4816, 4917, 4505, 4183, 2681, 4537, 633, 1321, 1977, 282, 4500, 1952, 4695, 4820, 1540, 4617, 3495, 267, 265, 2287, 633, 628, 3112, 2823, 1353, 3671, 3477, 4676, 3647, 299, 3808, 3771, 3340, 3274, 3251, 4820, 1679, 1681, 3778, 3518, 4902, 2841, 3559, 1931, 4422, 434, 1761, 3454, 4659, 3057, 905, 2086, 1043, 1272, 1260, 1725, 2384, 3550, 1823, 1969, 451, 371, 2602, 2535, 4911, 4750, 990, 3331, 2686, 1282, 543, 140, 2998, 2407, 1370, 3810, 2252, 4582, 3781, 2964, 496, 20, 4087, 4115, 1039, 1428, 2939, 1761, 3493, 631, 2169, 2392, 2801, 1596, 4790, 2767, 278, 4829, 4345, 4757, 725, 3027, 1999, 390, 379, 4541, 2755, 2907, 3978, 4536, 2814, 2381, 2443, 768, 2912, 3739, 2283, 57, 597, 3745, 1345, 1223, 4626, 1681, 394, 2843, 4687, 1248, 21, 4988, 30, 1315, 1300, 4652, 3548, 3181, 8, 4917, 3888, 3444, 936, 1666, 4038, 3825, 4977, 3868, 643, 3352, 2828, 2293, 4919, 675, 1520, 1439, 4979, 4807, 1547, 1265, 2687, 1474, 1931, 567, 768, 1223, 2021, 664, 1296, 2361, 1368, 4545, 1929, 1151, 3597, 2320, 458, 2967, 1208, 4533, 1600, 1762, 4799, 3081, 1123, 162, 2494, 4849, 140, 2697, 1490, 3699, 3261, 4499, 2337, 1827, 1921, 4324, 4274, 30, 3925, 1605, 954, 4607, 2340, 2532, 1591, 4365, 2352, 3728, 3314, 3787, 1743, 4585, 848, 2989, 1947, 2029, 3380, 1703, 3222, 4305, 3277, 3718, 3615, 195, 2853, 4657, 3804, 3366, 1635, 1783, 44, 2998, 3939, 4685, 4504, 2853, 3075, 411, 474, 1361, 2127, 1887, 1716, 2327, 2526, 3090, 2529, 2314, 2343, 366, 3419, 3303, 3491, 3931, 2944, 3512, 2624, 3871, 3407, 3068, 3239, 2702, 1516, 2835, 2333, 3246, 242, 2923, 833, 1565, 1857, 327, 2609, 1643, 4427, 2578, 503, 849, 2717, 3920, 2465, 4658, 3243, 1051, 4804, 721, 4824, 4390, 2140, 2709, 216, 806, 4978, 121, 3220, 2132, 1565, 1594, 4869, 4072, 3351, 3682, 3987, 1273, 104, 1272, 2324, 4744, 265, 3512, 669, 566, 4514, 4794, 2418, 2497, 1989, 2728, 3360, 4451, 2037, 320, 3354, 4204, 3480, 1842, 1729, 4032, 3762, 844, 2252, 2443, 697, 3783, 2478, 4625, 3265, 1168, 2581, 3552, 3182, 3156, 15, 2839, 227, 1085, 4491, 2220, 1408, 3286, 462, 2681, 2909, 1709, 2126, 3702, 4062, 3317, 1424, 3712, 958, 3345, 2386, 2555, 3288, 3678, 2593, 2711, 2304, 2539, 4648, 4605, 2153, 4767, 3873, 1684, 1742, 2023, 3943, 1632, 2074, 3816, 1482, 4942, 3726, 2520, 83, 3502, 3380, 4431, 4884, 3779, 4858, 4476, 4409, 4924, 2758, 401, 4561, 1144, 3302, 3374, 965, 3827, 2716, 2750, 228, 4300, 4894, 147, 4148, 2986, 3963, 2858, 3801, 1311, 2471, 3538, 1768, 3154, 1487, 3901, 753, 2133, 1555, 2467, 3735, 2459, 216, 4864, 803, 4827, 1641, 1531, 3361, 3314, 1438, 3364, 3485, 3876, 4566, 420, 3421, 2079, 1557, 429, 4788, 3242, 306, 4319, 3959, 2811, 2248, 4119, 936, 1111, 2048, 1741, 1404, 3556, 3060, 3055, 942, 3737, 513, 3136, 657, 3883, 3443, 193, 1879, 579, 478, 2279, 1292, 1269, 4272, 3843, 3142, 770, 1203, 4809, 544, 4248, 4483, 4146, 352, 3808, 250, 4075, 2838, 372, 3248, 1639, 3679, 405, 261, 646, 469, 1183, 1467, 95, 377, 2661, 1696, 4463, 4133, 4190, 1984, 598, 4271, 1287, 1164, 1186, 51, 4611, 1118, 3033, 3989, 3223, 3494, 1089, 9, 1238, 3930, 1273, 1409, 4758, 3656, 213, 155, 1434, 3286, 3862, 2714, 2216, 596, 2100, 1809, 1512, 349, 804, 1861, 1968, 3537, 2732, 2412, 2625, 90, 11, 1993, 3347, 4664, 2306, 231, 2050, 1141, 1844, 859, 680, 2994, 1309, 4274, 1853, 1462, 821, 2928, 685, 3506, 748, 1589, 4885, 414, 3061, 1649, 2048, 4659, 2892, 1040, 2196, 2069, 3496, 4924, 1386, 4398, 96, 3418, 2217, 3495, 2736, 4120, 3306, 4214, 3302, 3971, 2029, 2305, 1854, 4026, 3944, 2597, 2890, 1804, 1629, 571, 1305, 773, 704, 1520, 4440, 3002, 1728, 4777, 4730, 645, 1765, 4, 3687, 4975, 2862, 4557, 1444, 74, 782, 2339, 4122, 3395, 4494, 834, 3746, 3356, 2770, 4326, 4976, 948, 2673, 2712, 1522, 1298, 3043, 1134, 3210, 618, 1754, 4855, 2219, 1002, 3449, 4754, 4358, 2527, 2640, 182, 2788, 4305, 1148, 475, 4817, 2070, 2920, 441, 4510, 683, 4872, 738, 4874, 4490, 3265, 3565, 3091, 2877, 2498, 3876, 840, 2738, 4307, 1855, 1723, 2488, 3008, 2830, 2652, 1530, 1271, 3540, 61, 20, 4073, 1095, 1742, 1246, 2548, 1770, 2724, 3473, 4700, 2542, 3641, 1303, 4334, 3690, 1190, 2599, 1937, 3735, 3384, 4717, 974, 1099, 3023, 1118, 4825, 289, 1975, 2465, 3737, 4864, 3131, 2583, 115, 4952, 4656, 1321, 3533, 1051, 2532, 3133, 4595, 819, 3351, 437, 526, 2341, 501, 875, 4362, 1276, 2703, 472, 350, 4373, 160, 3713, 75, 3789, 3697, 2518, 1681, 1525, 3143, 445, 44, 2253, 3077, 1042, 2896, 2435, 4515, 1782, 873, 764, 4474, 1305, 3611, 2489, 4451, 16, 3801, 3703, 3928, 4837, 2366, 3386, 3718, 2083, 320, 3595, 587, 1515, 3045, 293, 2793, 4546, 845, 1160, 1369, 4754, 1457, 2071, 4248, 897, 3888, 1637, 928, 2227, 1792, 259, 3184, 3147, 2925, 4827, 1787, 746, 2685, 2951, 3717, 1709, 1532, 26, 1004, 3316, 630, 3730, 1574, 2535, 2379, 2172, 2141, 3116, 4269, 4967, 1997, 2663, 1980, 4515, 2762, 3930, 2811, 2615, 2383, 2415, 2670, 1214, 618, 4374, 2620, 1607, 1007, 4513, 1827, 3890, 1194, 2681, 4772, 2025, 1195, 3374, 138, 135, 4457, 1461, 1362, 3204, 783, 2207, 1888, 4771, 2590, 3143, 217, 4884, 4047, 1123, 3915, 2572, 2234, 3607, 3958, 2297, 1042, 2632, 1498, 1007, 686, 312, 3697, 2286, 3123, 794, 472, 2302, 1228, 803, 4029, 1416, 3693, 4897, 444, 906, 4990, 3213, 1216, 1881, 2982, 2840, 1186, 2434, 3322, 2883, 416, 4646, 2197, 3353, 295, 2326, 3418, 1935, 2284, 2507, 4352, 940, 122, 967, 68, 3335, 2214, 3063, 1088, 2293, 4821, 1557, 2854, 440, 4087, 965, 3247, 3554, 3820, 2952, 912, 9, 3458, 1560, 1827, 1109, 2560, 3539, 2872, 3061, 960, 2919, 332, 3322, 1709, 887, 2465, 3553, 393, 3771, 2852, 1110, 2070, 4726, 1529, 4215, 2452, 43, 2539, 3425, 2711, 1747, 3086, 2245, 2483, 4415, 613, 4, 4671, 3894, 240, 596, 2719, 2739, 1513, 2011, 2967, 2136, 2851, 1954, 1579, 274, 4303, 2083, 4801, 3001, 2126, 4633, 3932, 1468, 2712, 2711, 3747, 397, 312, 1340, 322, 4079, 3792, 4465, 4304, 3676, 1527, 1387, 2948, 4755, 2478, 4364, 4129, 1123, 1032, 619, 4454, 850, 1996, 617, 713, 1, 679, 2693, 406, 2693, 200, 180, 1836, 2521, 2032, 3571, 261, 997, 3239, 4927, 2645, 4061, 2071, 3413, 4906, 4944, 2391, 4255, 1799, 398, 1303, 2211, 4783, 1901, 572, 2985, 515, 1268, 4051, 4275, 1587, 1009, 184, 187, 3116, 2100, 4398, 4887, 765, 1658, 3272, 2396, 2321, 2756, 3974, 4513, 4408, 123, 4694, 2712, 752, 342, 3467, 714, 404, 539, 4182, 4650, 3597, 2406, 918, 1706, 677, 1381, 3127, 4756, 4140, 409, 773, 2855, 4869, 378, 4631, 2587, 3321, 165, 4887, 3742, 4138, 3519, 4179, 4442, 598, 2084, 4296, 1120, 4438, 4592, 2253, 4056, 2396, 730, 233, 2235, 4491, 2542, 3600, 3047, 3428, 1522, 3665, 2528, 1005, 2638, 2231, 1070, 4998, 1785, 4581, 2621, 1234, 2763, 1618, 314, 4551, 240, 4593, 1064, 3813, 874, 2727, 1959, 4366, 2334, 147, 3259, 4675, 4391, 284, 3308, 1821, 1809, 1199, 4552, 4521, 4837, 660, 3367, 1690, 1305, 1364, 3140, 1048, 370, 3472, 4647, 2234, 1574, 522, 2342, 1914, 3511, 2408, 384, 4202, 3362, 3243, 1858, 4355, 780, 3050, 764, 246, 4521, 1568, 1432, 2303, 848, 639, 4746, 4834, 4854, 4379, 3813, 3715, 1301, 3048, 3152, 2759, 4152, 3298, 884, 3314, 2071, 895, 3070, 395, 2607, 3737, 2292, 3524, 4691, 3470, 3122, 3276, 2388, 4624, 681, 4474, 703, 3328, 3937, 336, 4984, 1136, 2293, 4247, 1288, 1017, 4300, 1595, 263, 1404, 2205, 1157, 2710, 257, 757, 4158, 1938, 2380, 3969, 2606, 826, 2282, 2966, 526, 4518, 2333, 1820, 812, 627, 2287, 1906, 2778, 226, 2164, 3408, 2489, 2186, 919, 211, 3302, 3470, 4594, 2842, 4535, 3605, 1131, 3903, 2859, 4711, 3194, 4136, 1267, 3818, 4626, 433, 4496, 3331, 4222, 2723, 252, 4067, 245, 294, 3852, 3714, 1736, 2460, 3895, 1624, 849, 4458, 1088, 3818, 1882, 2165, 4749, 4854, 616, 3043, 1074, 4867, 4055, 4916, 2183, 4189, 1359, 4965, 1578, 1177, 718, 2238, 404, 404, 1468, 3957, 401, 4875, 2721, 2312, 3518, 4331, 2045, 2273, 1790, 4615, 2057, 4701, 535, 1339, 1732, 4596, 548, 462, 2135, 1089, 4697, 2774, 2048, 2053, 1475, 2360, 1129, 3413, 1433, 965, 2697, 117, 1167, 2315, 3460, 4122, 571, 3314, 873, 2334, 708, 939, 4199, 1583, 3257, 792, 2145, 695, 2514, 3307, 743, 784, 814, 1541, 4948, 3787, 4314, 595, 1453, 3788, 2795, 3221, 611, 774, 2073, 4638, 633, 1202, 273, 1224, 4003, 2426, 4193, 497, 1712, 1233, 368, 2684, 87, 3598, 4505, 1477, 3905, 1378, 3948, 476, 3734, 1639, 2222, 4451, 798, 2892, 4186, 4452, 1692, 2208, 4290, 1244, 2375, 2418, 3527, 4021, 4336, 613, 4926, 3462, 3764, 48, 924, 780, 3127, 16, 1084, 1175, 1545, 1265, 3138, 4541, 2184, 561, 1522, 4048, 64, 3929, 327, 2483, 3321, 2369, 2199, 4164, 2995, 4707, 1839, 3519, 918, 2206, 3332, 2706, 2442, 521, 767, 1000, 1008, 4475, 276, 1329, 337, 4418, 2125, 1518, 1366, 4767, 3371, 589, 2624, 4141, 1678, 1491, 137, 3224, 97, 2161, 114, 831, 2972, 4717, 2233, 3721, 4621, 2017, 2475, 3752, 743, 321, 1559, 3279, 3690, 4444, 3542, 2605, 1349, 1018, 2305, 4496, 1479, 1529, 1920, 4471, 2553, 3883, 512, 4645, 3323, 1786, 3944, 1946, 2668, 443, 4442, 1078, 3992, 3859, 1429, 4964, 2077, 4290, 4877, 1941, 1099, 3756, 2443, 181, 415, 4415, 3253, 3968, 2298, 3668, 352, 4495, 4034, 1212, 3832, 4856, 4869, 3106, 3376, 3633, 3310, 2017, 1476, 4327, 2833, 3183, 157, 17, 2911, 3006, 4317, 3614, 4473, 4917, 2606, 555, 2833, 1188, 4303, 277, 4245, 3651, 4318, 4923, 4924, 1965, 94, 1469, 2415, 2682, 1397, 3386, 230, 3477, 3825, 434, 3116, 1995, 1259, 2389, 1562, 3016, 2522, 193, 576, 4636, 4310, 4317, 4767, 3521, 1552, 140, 3570, 3467, 1405, 3090, 942, 564, 4800, 3981, 4865, 667, 3459, 4916, 4577, 1921, 4906, 1492, 1432, 1289, 1397, 284, 3713, 3747, 163, 2419, 4983, 948, 4682, 689, 2147, 4515, 1933, 542, 2225, 4410, 3830, 1160, 100, 1119, 4369, 2231, 473, 3331, 4776, 2471, 2976, 3091, 2740, 2763, 1597, 4700, 2860, 4410, 2165, 3369, 2034, 1840, 2555, 4088, 1712, 2169, 3303, 683, 1916, 4023, 2933, 3341, 4604, 3312, 4117, 1788, 649, 2800, 3377, 242, 3367, 2316, 2612, 4725, 1657, 4042, 1908, 2355, 2954, 2028, 3185, 4256, 854, 4082, 668, 4999, 323, 3720, 2146, 3703, 674, 1338, 4008, 4152, 644, 1870, 3553, 2184, 2969, 4044, 2228, 751, 1972, 893, 1601, 4562, 3571, 2846, 2319, 1127, 4815, 1737, 1728, 3338, 1920, 3229, 2103, 2293, 2765, 2602, 4239, 176, 4333, 472, 3340, 3551, 1443, 4630, 3093, 1965, 659, 779, 1393, 4346, 1459, 1780, 2907, 3795, 3561, 200, 646, 2752, 3493, 1801, 3360, 2458, 1914, 3932, 4558, 1123, 2024, 4818, 4852, 2552, 1921, 2842, 4088, 727, 2576, 2288, 4934, 2737, 699, 326, 4723, 3759, 4412, 1321, 3713, 1322, 115, 4508, 3435, 2226, 2472, 1095, 3809, 4684, 4058, 500, 4812, 2669, 604, 3612, 2748, 1283, 3332, 3690, 2200, 4654, 2765, 1722, 4840, 2771, 1654, 3686, 2266, 2973, 4363, 3799, 3136, 2894, 111, 440, 565, 1211, 2640, 2833, 2114, 179, 3829, 3714, 2564, 2380, 690, 1851, 2092, 3571, 2401, 2724, 1789, 4392, 1866, 1320, 2821, 1920, 2903, 1917, 498, 2023, 529, 3734, 1950, 1112, 3653, 1488, 2246, 1723, 1472, 3340, 2108, 4807, 1149, 4563, 1428, 4804, 3823, 2768, 3112, 2636, 538, 625, 483, 432, 1312, 2457, 3953, 620, 347, 565, 1495, 2494, 3091, 1457, 3413, 1808, 4724, 1726, 2355, 3885, 2913, 653, 3928, 3874, 2844, 905, 3764, 705, 4058, 2985, 3946, 4855, 2577, 613, 209, 4207, 3192, 4573, 332, 2779, 669, 3928, 3604, 3157, 2037, 19, 608, 3354, 1790, 568, 1908, 2092, 460, 3743, 4226, 3691, 803, 4940, 1053, 2619, 742, 1834, 606, 2878, 2202, 1121, 2314, 631, 2883, 2396, 4405, 3644, 4583, 4852, 3593, 2322, 2368, 4616, 1693, 2703, 3232, 2498, 4762, 3613, 785, 3958, 3744, 598, 49, 2788, 3046, 4526, 3842, 955, 1583, 1520, 4183, 4002, 593, 2120, 3501, 4695, 4245, 593, 1119, 786, 3355, 3998, 191, 4154, 1250, 1916, 1240, 3011, 1965, 2853, 968, 4945, 1745, 1172, 1220, 4237, 1655, 2044, 2258, 1038, 3355, 2415, 803, 4353, 2939, 1254, 3415, 656, 2429, 4586, 2912, 719, 4592, 2464, 3034, 4839, 1783, 659, 4505, 2613, 1988, 4718, 1150, 71, 2143, 4653, 2397, 3877, 3617, 1816, 1047, 65, 2778, 1690, 1164, 1729, 3329, 3499, 2985, 4393, 753, 4171, 3209, 4733, 4336, 2105, 2313, 2047, 2114, 3084, 4982, 3277, 2352, 61, 662, 498, 4214, 3837, 3357, 2059, 2469, 2553, 1956, 4824, 3302, 598, 3450, 1129, 2542, 2072, 536, 1393, 2010, 2278, 3906, 1984, 935, 822, 2170, 3487, 1587, 3990, 2556, 521, 3150, 2921, 495, 1809, 436, 2689, 1587, 2983, 1701, 945, 1165, 2875, 4641, 4427, 150, 1922, 2039, 960, 4139, 2052, 4523, 3252, 4841, 408, 1790, 2286, 1598, 4539, 3839, 2944, 4598, 2806, 2765, 1279, 2807, 3047, 1036, 2188, 4023, 3980, 4816, 4618, 1238, 1974, 1043, 1699, 1894, 4579, 652, 915, 3316, 692, 726, 3076, 1794, 4611, 3202, 498, 2560, 2481, 4409, 4780, 1772, 3646, 1750, 3676, 3047, 2572, 4891, 4460, 3044, 3855, 800, 1141, 4269, 4297, 2132, 643, 599, 3600, 310, 1395, 1368, 3408, 2681, 2990, 3558, 536, 644, 3325, 4128, 516, 1236, 770, 1193, 183, 2729, 656, 2097, 3905, 4273, 390, 4885, 203, 1557, 3853, 1763, 2856, 4218, 1795, 2537, 1806, 32, 972, 1022, 1508, 4843, 3870, 2394, 1211, 4003, 1232, 3372, 1303, 3911, 4898, 1370, 3516, 4833, 3955, 1788, 4505, 1449, 1237, 2553, 877, 872, 35, 2970, 4564, 4512, 1102, 4181, 1229, 1525, 4059, 4554, 3889, 1865, 2439, 4390, 1949, 3854, 4363, 182, 4310, 2501, 1905, 2183, 4824, 3576, 1593, 4260, 2314, 4656, 1096, 460, 4947, 4831, 1313, 2923, 3726, 3857, 4474, 1541, 1443, 3415, 2367, 3195, 2916, 513, 306, 453, 457, 1332, 2141, 309, 251, 691, 455, 3216, 4044, 1267, 87, 4069, 31, 3130, 4215, 1717, 103, 636, 1698, 752, 4211, 89, 4713, 3369, 1924, 3499, 1686, 875, 3597, 276, 3784, 2381, 4326, 2267, 4390, 170, 2800, 4517, 2740, 2322, 4494, 1373, 1498, 3027, 2874, 2757, 1547, 1336, 1471, 1992, 3121, 2449, 726, 2088, 911, 4582, 3658, 333, 3547, 1785, 4799, 2703, 21, 411, 1807, 2123, 761, 2590, 4243, 4741, 4897, 1410, 253, 1246, 672, 1136, 3506, 4459, 1869, 1588, 582, 579, 4908, 1679, 3827, 1898, 3390, 3515, 3692, 3554, 3413, 4262, 4985, 4882, 4525, 1767, 1312, 4073, 2665, 1992, 1110, 1150, 1612, 4704, 4677, 1628, 4477, 292, 3256, 899, 2106, 410, 2657, 227, 2673, 760, 4080, 3472, 1751, 3267, 3585, 4281, 1749, 2606, 798, 2523, 809, 2181, 251, 4778, 4588, 2974, 3943, 4317, 3434, 4405, 4794, 4260, 4358, 3901, 4812, 3184, 2369, 3960, 2729, 3575, 528, 1664, 4107, 2423, 2464, 442, 1189, 3677, 862, 4916, 495, 3436, 4172, 612, 3128, 1876, 1239, 2300, 3751, 3349, 473, 2180, 3411, 548, 1983, 2542, 2127, 2950, 130, 3594, 2391, 3812, 4199, 2565, 1265, 1961, 3223, 65, 3578, 4744, 1714, 3744, 232, 277, 618, 4248, 1620, 2426, 4164, 4350, 3593, 2356, 2890, 3758, 2828, 4155, 3789, 4827, 3726, 3375, 20, 3814, 4604, 2017, 2461, 4351, 1372, 2482, 1691, 1517, 2300, 4210, 3963, 4891, 2646, 4546, 1737, 89, 2338, 3668, 3234, 3957, 2169, 1192, 1029, 4079, 1159, 2752, 2590, 4360, 4672, 3581, 503, 4046, 3116, 210, 3186, 841, 2630, 1528, 3148, 835, 248, 2251, 175, 1965, 3928, 1416, 3298, 3107, 237, 2271, 3830, 1874, 871, 313, 1673, 2024, 1079, 474, 1930, 2369, 3000, 4769, 4300, 418, 3572, 4977, 2186, 1728, 682, 1270, 3884, 2977, 2667, 117, 643, 1461, 2358, 1299, 1589, 1969, 4374, 1861, 2736, 983, 1101, 658, 1929, 1959, 2802, 166, 129, 2011, 2983, 2796, 1585, 3299, 4242, 2735, 3143, 1018, 138, 2593, 3128, 3990, 3080, 3707, 3958, 4644, 4430, 3213, 1754, 2685, 1182, 4672, 1023, 2517, 2580, 3706, 4804, 2136, 2388, 343, 3103, 4443, 1496, 2625, 3511, 2989, 921, 4127, 1666, 4310, 2992, 4367, 633, 1287, 3277, 2440, 771, 882, 2527, 1629, 2360, 1461, 2928, 2676, 2997, 4906, 4090, 187, 2650, 3376, 4517, 3229, 1108, 3392, 1997, 3136, 4079, 1677, 2497, 3388, 4413, 2621, 1132, 2975, 4313, 4991, 3831, 4052, 4434, 3484, 4962, 992, 1194, 305, 262, 343, 4031, 1961, 1869, 1491, 1884, 3066, 3791, 3529, 970, 2021, 4584, 1801, 902, 1251, 2390, 2911, 2153, 3996, 1687, 3040, 577, 2428, 4495, 1584, 2928, 2012, 3163, 211, 2062, 1319, 667, 3808, 806, 512, 4259, 967, 1868, 2714, 4011, 2305, 2853, 2438, 1936, 1208, 3755, 3330, 771, 4783, 862, 2541, 1537, 31, 2950, 4772, 462, 1165, 3879, 2973, 1361, 1093, 1005, 2968, 1807, 4969, 1793, 881, 61, 1303, 4355, 1277, 347, 1379, 4687, 1426, 2053, 269, 2328, 3378, 3377, 1043, 3483, 2406, 2648, 411, 3074, 3034, 1125, 3711, 1495, 2155, 2100, 4285, 4629, 1958, 4863, 2634, 1693, 2015, 3586, 4446, 2088, 4523, 4831, 3722, 4207, 3549, 4764, 3259, 4686, 2666, 751, 2980, 2308, 598, 2923, 1026, 932, 1523, 3679, 401, 212, 597, 3517, 4675, 2172, 3847, 1588, 1952, 1906, 3896, 3284, 4671, 453, 3215, 4686, 1280, 348, 343, 3542, 576, 1580, 3570, 694, 2654, 707, 1037, 1811, 3700, 3071, 444, 4676, 3493, 4212, 3227, 3800, 4459, 71, 687, 2901, 1862, 3189, 2739, 884, 4675, 4445, 1121, 865, 701, 3234, 4819, 4393, 1867, 350, 4120, 4840, 664, 681, 1738, 4656, 2708, 4506, 1237, 4125, 3090, 1760, 4342, 4893, 4768, 4108, 452, 778, 2671, 4658, 916, 3872, 1013, 806, 2754, 1892, 1422, 2449, 814, 2823, 1319, 1636, 1554, 2269, 1686, 2627, 4837, 676, 4342, 1830, 2334, 4616, 2797, 4205, 25, 3796, 4289, 1002, 1118, 2986, 2695, 2406, 1000, 4996, 1604, 1508, 1789, 2708, 4476, 2590, 3446, 4512, 1893, 3930, 3346, 1232, 4776, 1529, 3893, 2220, 1195, 1997, 1596, 4822, 1721, 4392, 3901, 3979, 4543, 703, 3918, 3537, 1242, 4632, 3048, 206, 2938, 3576, 4384, 4507, 4918, 1868, 1073, 2281, 793, 3028, 3295, 1618, 122, 4316, 1847, 4152, 3781, 3505, 2740, 1481, 3565, 3765, 3978, 272, 1335, 2744, 950, 880, 1360, 2189, 615, 3527, 872, 4623, 2474, 3100, 4751, 3536, 1187, 3599, 1372, 281, 4372, 4997, 1478, 959, 3727, 532, 4287, 742, 4743, 3783, 1737, 3338, 2639, 2414, 511, 3120, 3313, 1555, 815, 1973, 4419, 757, 2732, 2659, 2392, 3755, 2992, 463, 915, 4883, 2441, 3861, 4413, 4551, 2972, 4875, 1727, 3994, 2985, 2591, 2984, 3495, 2754, 4200, 860, 3668, 2948, 9, 4628, 2370, 4188, 2923, 4604, 4103, 683, 3605, 587, 3272, 965, 2854, 3280, 2205, 196, 4939, 2466, 4674, 225, 2207, 4990, 4387, 1932, 3194, 2587, 2875, 1419, 4446, 3093, 3022, 3196, 3901, 2507, 3096, 4080, 4179, 1051, 4741, 188, 1708, 1862, 3830, 3642, 1399, 2931, 4847, 2477, 353, 1864, 1798, 2745, 4237, 4897, 67, 3863, 2496, 1409, 3345, 3480, 4196, 683, 1917, 2594, 4083, 782, 2399, 3611, 8, 220, 1539, 3847, 1662, 3001, 3850, 673, 2099, 1014, 3753, 4601, 4256, 2310, 729, 4526, 3753, 185, 2618, 1621, 1212, 2688, 3351, 3749, 3093, 3135, 4052, 4639, 4308, 3466, 1015, 4503, 2613, 2742, 4046, 3930, 3153, 2014, 296, 1767, 4201, 2161, 4000, 4681, 1245, 4320, 3488, 2386, 4559, 225, 3865, 3033, 487, 1904, 4850, 2924, 1938, 327, 642, 4378, 1337, 3466, 2118, 1979, 2997, 4563, 4024, 1607, 4020, 759, 1691, 4296, 2075, 126, 3675, 2835, 3230, 3745, 556, 4468, 454, 3883, 4312, 882, 1020, 4351, 156, 3311, 795, 2113, 1827, 2121, 2063, 304, 1648, 621, 3689, 3591, 3424, 1459, 2174, 74, 3670, 4827, 3084, 2954, 2828, 2812, 1387, 4187, 2965, 1839, 3234, 691, 3193, 2772, 2771, 651, 3942, 2990, 2728, 1635, 432, 310, 4336, 2208, 3151, 3058, 624, 4551, 1546, 646, 945, 2973, 2881, 3667, 2521, 2611, 821, 1421, 4767, 2287, 3845, 4817, 1087, 2357, 822, 2263, 4325, 2781, 982, 284, 1543, 2579, 1037, 2080, 2543, 3428, 2940, 2955, 1450, 3661, 3463, 2913, 1049, 2912, 3897, 2779, 1910, 3218, 2245, 4588, 1826, 2941, 187, 2296, 1995, 1686, 303, 3571, 743, 3833, 242, 2272, 900, 829, 2019, 3712, 1418, 2400, 4580, 551, 3869, 1435, 4003, 129, 2609, 80, 933, 3703, 1681, 703, 3999, 2156, 2936, 4223, 1630, 1650, 1728, 4322, 4368, 2186, 137, 4911, 4514, 3167, 3489, 3386, 3689, 4758, 2476, 1316, 4619, 264, 3829, 2656, 286, 812, 4093, 4030, 3154, 4211, 407, 1445, 2329, 2504, 1506, 3121, 2888, 2962, 3186, 1081, 4377, 207, 1240, 709, 1113, 3153, 818, 519, 1904, 3673, 2802, 1760, 4042, 740, 2102, 3602, 2828, 1486, 2633, 1814, 2774, 4150, 480, 1907, 4942, 3458, 3137, 213, 4316, 4292, 433, 2476, 3092, 3319, 2120, 3355, 855, 809, 1502, 4862, 4802, 2722, 1463, 4767, 3263, 4563, 806, 932, 3068, 3909, 4615, 1285, 390, 218, 4772, 22, 62, 4264, 1258, 54, 1501, 4674, 2749, 3934, 4557, 4962, 2620, 1678, 3510, 4418, 2940, 671, 2702, 968, 2600, 4427, 1676, 1383, 1190, 2190, 4229, 3146, 1620, 4309, 2571, 585, 302, 3112, 309, 4523, 2233, 3360, 117, 3317, 3767, 3218, 4961, 3572, 4823, 4272, 2881, 1876, 4812, 4693, 2581, 1915, 2244, 1561, 1276, 4067, 3589, 4613, 1593, 3561, 2955, 1397, 2328, 2962, 526, 4571, 4668, 2627, 1995, 4788, 3309, 2866, 1916, 4372, 1828, 161, 3548, 1393, 3078, 2749, 4924, 1157, 4556, 3450, 4545, 3545, 1171, 2596, 4122, 1960, 682, 3831, 4994, 2140, 4466, 3486, 2653, 3710, 4057, 4515, 2912, 1096, 3747, 2473, 2709, 1466, 2826, 4302, 3554, 3061, 2281, 1680, 1068, 1519, 2474, 4112, 3291, 858, 659, 3836, 1996, 1118, 68, 4729, 1024, 3399, 2501, 1012, 3229, 1851, 512, 779, 1576, 4590, 965, 4616, 4015, 1958, 4589, 3290, 3472, 4967, 2121, 4370, 4495, 4302, 1391, 3664, 2102, 1727, 4718, 2030, 3709, 3527, 674, 2735, 467, 1572, 3264, 1458, 307, 1881, 803, 1583, 3140, 4422, 4915, 1424, 360, 2781, 711, 4501, 2085, 301, 4623, 2360, 3929, 144, 4726, 4812, 1424, 957, 4699, 4760, 2319, 1421, 1531, 3346, 1967, 2445, 1647, 4801, 586, 952, 3383, 4856, 1577, 2756, 3819, 3878, 4354, 4349, 93, 4891, 2163, 1559, 12, 2520, 1776, 3534, 4194, 3045, 2730, 180, 3521, 4558, 441, 1771, 3119, 2980, 3809, 3439, 2624, 3315, 368, 3417, 33, 1400, 4847, 3433, 3267, 1419, 296, 4612, 1633, 3757, 2453, 2394, 486, 3144, 1857, 2363, 691, 2234, 4950, 1869, 4989, 4869, 4447, 4755, 846, 4951, 670, 4000, 2412, 491, 3481, 2563, 2698, 802, 3271, 3643, 2022, 4849, 1865, 777, 52, 672, 1276, 257, 1408, 4646, 956, 2601, 2579, 4210, 3844, 1288, 2402, 3842, 1023, 659, 2688, 4122, 2892, 2443, 461, 3886, 3652, 4174, 3200, 3501, 806, 4078, 2965, 2862, 3405, 2079, 4660, 4671, 2143, 475, 4226, 1252, 3510, 1133, 3946, 2800, 4307, 339, 897, 182, 1967, 915, 2737, 4211, 4268, 3089, 4419, 2088, 2433, 531, 3306, 2866, 2220, 1512, 702, 4229, 1042, 2126, 3442, 2922, 2613, 562, 3897, 3686, 4931, 3733, 2318, 4148, 2830, 262, 278, 4609, 1381, 416, 1016, 2529, 109, 1412, 1865, 2519, 2539, 4868, 1600, 1436, 2332, 3644, 3299, 2481, 2236, 2479, 1020, 1363, 4318, 3227, 2262, 3677, 2179, 4877, 1495, 818, 827, 4638, 4669, 3600, 4480, 3475, 2619, 1770, 1594, 4123, 4145, 2438, 3698, 1029, 2313, 3974, 2482, 360, 2841, 2687, 931, 3015, 769, 2587, 3453, 502, 2762, 1882, 3558, 982, 953, 4501, 1994, 2515, 4706, 1228, 4790, 582, 1157, 1682, 2801, 1649, 2052, 1583, 1079, 2885, 1143, 2352, 4301, 2553, 2758, 4001, 572, 2469, 4739, 2985, 1540, 3862, 2680, 1154, 3791, 3221, 3495, 2848, 504, 2329, 656, 1971, 1679, 1647, 4685, 498, 4714, 681, 1862, 4918, 1949, 4635, 3063, 19, 2759, 2109, 800, 3299, 4912, 3100, 2187, 3303, 2961, 2938, 369, 3071, 2682, 4425, 3113, 2521, 3321, 1865, 2526, 2409, 3826, 3298, 4718, 2506, 3487, 4782, 2848, 1415, 129, 1392, 4685, 2036, 1582, 4131, 1319, 1196, 2945, 3282, 1899, 1775, 1187, 2186, 4007, 506, 4091, 4599, 4470, 869, 4302, 1387, 2956, 4232, 1719, 1823, 1520, 4341, 1660, 4946, 4565, 2465, 2361, 2624, 3693, 2084, 3530, 1197, 3645, 1386, 4513, 2550, 1612, 1530, 2141, 1505, 1237, 2241, 811, 2283, 3323, 3559, 1833, 3525, 2278, 4458, 1090, 154, 4695, 4460, 1116, 1391, 2582, 4246, 1895, 3329, 265, 1938, 2830, 4749, 1703, 4202, 2122, 131, 950, 1661, 2510, 4854, 3353, 2761, 2190, 1921, 3593, 4312, 514, 2952, 4506, 849, 2438, 3908, 674, 2847, 1450, 1837, 27, 3715, 1504, 1698, 3403, 1665, 987, 4730, 863, 4407, 2667, 3881, 4396, 3142, 2686, 1827, 584, 3670, 1942, 669, 2773, 220, 374, 4749, 353, 4695, 1727, 2656, 3619, 2302, 3939, 3943, 4231, 117, 3725, 2126, 2327, 2756, 1353, 3959, 4928})
	fmt.Println(v)
}

func TestTindLengthOfLCIS(t *testing.T) {
	v := findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4})

	fmt.Println(v)
}

func TestCommonChars(t *testing.T) {
	v := commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"})

	fmt.Println(v)
}

func TestAddToArrayForm(t *testing.T) {
	v := addToArrayForm([]int{2, 1, 5}, 806)
	fmt.Println(v)
}

func TestLargeGroupPositions(t *testing.T) {
	v := largeGroupPositions("abbxxxxzzy")
	fmt.Println(v)
}

func TestFlipAndInvertImage(t *testing.T) {
	v := flipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}})
	fmt.Println(v)
}

func TestPivotIndex(t *testing.T) {
	v := pivotIndex([]int{-1,-1,-1,-1,0, 1})
	fmt.Println(v)
}
