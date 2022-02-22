package pipe

// data created by
// source ../script/staging_credentials && seed_data=1  go test -v --run=Seed
// in /uof/api directory
var marketGroupsJSON = `{"10_min":[593,574,734,577,591,576,578,573,107,105,595,575,592,594,106,572],"15_min":[565,103,102,569,568,590,566,589,571,588,570,586,104,587,567],"180s":[385,393,391,384,392,390,388,382,381,389,386,387,383],"1st_half":[69,62,176,486,150,68,602,489,482,157,402,178,75,179,618,74,619,70,905,151,483,180,73,181,173,488,81,61,154,484,72,78,160,182,80,152,552,542,490,485,472,161,635,471,175,153,63,64,71,183,66,156,159,76,174,60,636,149,158,77,65,79,637,470,487,155,177],"2nd_half":[83,95,96,97,88,90,84,98,91,543,92,545,85,86,544,553,87,94,93],"2nd_half_incl_ot":[294,295,232,231,293],"3_innings":[284,285,743,745,744,286],"4.5_innings":[282,281,279,280,283],"5_innings":[274,275,276,278,277],"5_min":[110,584,582,108,599,597,583,580,579,596,109,598,600,581,585],"bomb":[631,632],"bookings":[150,593,148,157,139,136,143,151,599,145,597,137,591,142,590,154,140,596,589,160,152,161,588,147,153,144,598,600,156,159,141,595,149,158,138,586,885,592,594,587,155,146],"break":[509,507,510,508,505,506,511,504],"checkout":[380,379],"combo":[820,858,1056,819,859,547,546,865,889,426,861,541,888,860,818,540,36,543,854,742,425,35,78,542,37,855,863,545,890,891,424,862,134,265,544,112,857,184,864,856,429,292,79,1055,423],"corners":[176,584,164,165,602,565,582,178,574,179,171,163,577,162,170,569,180,181,568,173,583,576,172,580,168,579,169,566,182,571,578,570,175,573,183,884,581,174,166,167,575,585,567,601,177,572],"cup":[804],"cup_group":[809,810,848,808],"cup_ko":[812],"cup_tie":[5,2,3,4],"delivery":[363,362],"dismissal":[344,343],"end":[322,323,325,324],"frame":[502,509,503,514,501,507,513,510,515,500,508,505,499,516,512,506,511,504],"game":[247,520,213,248,246,245,209,215,210,214,250,211,212],"hit":[1052],"incl_ei":[266,260,257,267,255,739,741,259,261,262,254,253,742,263,251,256,265,264,258],"incl_ot":[612,610,227,909,849,405,219,225,226,228,290,904,404,617,229,882,616,230,224,903,292,122,291,223],"incl_ot_and_pen":[411,421,412,420,413,426,419,418,417,425,427,407,424,406,416,414,410,409,415,408,423,422],"incl_so":[340],"inning":[288,746,747,748,749,287,750,751],"innings":[350,344,343,606,605,346,345,348,607,608,349,347],"kills":[630,734,554,555,726,723,627,731,626,724,333,629,628,623,622],"league":[804,806,805,807],"leg":[376,393,391,392,380,379,375],"map":[556,727,557,620,728,398,335,554,555,726,735,732,733,752,334,725,396,397,731,395,333,730,729,558],"map_1st_half":[736],"map_incl_ot":[625,621,330,331,624,332,623,622],"matchday":[801,802,803,798,797,799,800],"misc":[6,603,289],"offsides":[886],"ot":[114,117,463,116,118,897,465,464,112,113,466,115],"ot_1st_half":[119,120,121],"over":[357,358,356,359,360,361],"pen_so":[124,131,129,125,130,123,126,128,127,135,133,134,132],"penalties":[887],"period":[447,454,449,446,452,455,444,445,460,457,462,453,451,443,450,459,529,458,456,448],"pitch":[1049,1050,1051],"player":[625,1053,1054,633,735,624],"player_props":[792,791,784,774,783,781,765,763,766,773,776,789,761,768,790,782,775,764,796,762,787,788,769,795,778,777,770,780,794,760,793,767,759,785,772,771,779,786],"points":[218,216,217],"progress":[732,733,725],"quarter":[302,757,756,303,758,235,755,301,1057,754,236,753,305,304],"quarter_incl_ot":[613,611,615,614],"rapid_market":[887,1049,1052,1053,1054,1050,886,1051,883,884,885],"regular_play":[385,33,31,12,234,27,820,164,430,367,165,475,858,476,34,819,192,28,148,431,193,467,893,604,859,139,48,19,13,439,241,199,900,879,737,136,171,547,384,546,143,20,41,40,899,495,49,55,865,18,889,1058,269,526,474,163,365,300,634,440,433,53,190,898,432,162,170,46,861,328,541,329,297,145,491,888,549,548,860,740,818,853,8,220,268,480,137,191,473,563,47,540,7,36,101,272,198,496,201,25,26,56,45,273,142,172,10,16,168,854,140,38,169,270,326,29,271,237,723,35,481,368,881,969,1,532,57,188,196,30,492,896,399,23,37,855,1059,100,341,298,51,435,478,863,434,327,189,890,52,891,382,147,442,862,551,58,238,15,550,852,9,381,880,738,144,54,401,902,857,184,32,438,525,850,724,11,386,39,141,892,494,14,864,166,186,59,400,167,138,901,609,856,194,366,429,383,306,50,1060,1055,239,21,479,314,437,146,851,436,601,441,342,493,24,477,195,968,187],"round":[337,631,632,630,633,338,627,626,629,628],"score":[282,340,447,69,124,131,350,62,351,612,302,110,33,114,621,31,12,129,322,323,234,376,309,27,83,610,95,411,125,757,227,430,367,528,756,613,895,206,34,5,130,288,454,303,266,207,192,421,260,312,28,746,68,96,316,412,502,281,611,431,330,747,193,402,274,284,467,449,294,257,604,620,295,75,894,357,247,503,285,48,19,13,758,520,439,241,199,358,213,879,737,849,514,405,20,267,41,420,117,446,413,748,219,255,495,331,377,313,49,123,55,618,97,739,248,103,74,619,18,1058,225,501,269,527,126,108,526,88,226,365,300,119,70,370,373,905,634,90,741,246,235,440,433,755,53,190,301,749,335,325,432,324,318,606,46,463,116,328,228,102,329,452,297,491,84,605,290,549,118,548,740,371,897,853,73,419,8,98,220,208,91,259,455,315,268,444,191,287,1057,352,754,563,47,7,245,445,904,261,101,272,198,310,378,513,460,262,496,201,497,25,26,56,218,279,128,45,346,254,273,418,743,750,81,253,404,61,417,356,127,10,617,465,16,745,236,270,326,216,29,271,2,345,237,72,368,515,135,881,457,1,80,532,354,334,498,500,57,188,263,196,30,427,492,533,552,896,399,23,229,1059,363,209,100,341,133,407,203,751,298,51,372,435,472,462,616,434,251,230,327,453,451,256,189,92,52,635,353,406,471,753,499,442,85,416,443,217,551,215,58,238,15,317,550,348,63,607,210,883,852,9,450,280,86,880,738,109,202,464,120,64,516,54,224,71,375,107,232,401,902,121,355,205,311,553,512,414,66,204,32,275,105,362,438,525,374,850,305,214,113,11,250,211,459,76,494,276,332,14,60,636,87,903,186,410,529,409,59,359,3,458,466,278,212,94,608,349,400,615,231,609,456,194,277,77,65,736,366,304,306,50,1060,319,104,415,637,122,360,239,21,314,470,106,132,264,347,744,4,437,408,291,258,286,851,448,436,422,115,441,342,493,24,614,93,195,283,187,293,223,361],"scorers":[909,893,900,40,899,898,38,969,882,39,892,901,968],"set":[309,528,206,1056,207,312,316,313,527,390,370,373,318,371,208,315,310,388,203,372,317,389,202,205,311,204,374,387,319],"structures":[556,727,557,728,398,752,396,397,730,729,558],"tiebreak":[895,894],"tries":[486,475,476,489,482,474,483,480,473,488,484,481,490,485,478,479,487,477],"visit":[377,378],"x_frames":[497,498,533],"x_overs":[351,352,354,353,355]}`