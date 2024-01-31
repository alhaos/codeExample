# GoCFX specification

## gpp

Conditions to be checked before posting the results.

1. Internal control must always be PRES (present). If the internal control is not PRES, do not post the results

2. If the sample has No Call, do not post the results

3. If the sample has more than 1 pathogen present, do not post the results

Results as “NEGATIVE” if neg and POSITIVE when Pos.

Please note all these pathogens are also reportable, if positive the respective department of health has to be notified.

The test codes and the names of pathogens

| Code | Name                   |
| ---- | ---------------------- |
| X014 | Shigella               |
| X357 | Entamoeba histolytica  |
| X024 | Campylobacter          |
| X010 | Salmonella             |
| X359 | Cryptosporidium        |
| X520 | Adenovirus 40/41       |
| X140 | C. difficile toxin A/B |
| X198 | Vibrio cholerae        |
| X032 | STEC stx1/stx2         |
| X385 | Rotavirus A            |
| X026 | E. coli O157           |
| X355 | Giardia                |
| X040 | ETEC LT/ST             |
| X371 | Norovirus GI/GII       |

## mup

CutOff based tests ("UU": 34.0, "MH": 33.0, "MG": 31.0, "UP": 34.0, "IC": 40.0)

1. Test result with N/A value reported as ND

## ABG

Excel files are submitted as input

1. Data rows in files 
   
   - column with index 0 is integer
   - column with index 4 is accession 10 digits
   - column with index 9 is ct float
   - column with index 10 is sd float

2. Three lines for each test; tests containing a different number of lines are rejected

3. cutOff map from tests 

```
var cutOffMap = map[string]testData{
    "Ba07319991_s1": {"Z957", 28, 1, "aac(6')-Ib-cr"},
    "Ba04646145_s1": {"Z959", 28, 1, "qnrS"},
    "Ba04646160_s1": {"Z961", 27, 1, "qnrA"},
    "Ba07319988_s1": {"Z967", 25, 1, "sul1"},
    "Ba07320003_s1": {"Z969", 27, 1, "sul2"},
    "Ba04646152_s1": {"Z971", 30, 1, "BlaKPC"},
    "Ba04646149_s1": {"Z977", 29, 1, "CTXM-G1"},
    "Ba04646142_s1": {"Z979", 29, 1, "CTXM-G2-bla"},
    "Ba04646134_s1": {"Z981", 25, 1, "BlaOKP-C"},
    "Ba04931076_s1": {"Z991", 25, 1, "BlaNDM-1"},
    "Ba04646155_s1": {"Z993", 27, 1, "BlaVIM"},
    "Ba04646131_s1": {"Z995", 27, 1, "IMP-1-CarbB"},
    "Ba04646158_s1": {"C018", 26, 1, "BlaMP"},
    "Ba04646135_s1": {"N111", 26, 1, "BlaCMY"},
    "Ba04646120_s1": {"T154", 28, 1, "DHA beta-lactamase"},
    "Ba04646126_s1": {"Z112", 27, 1, "FOX-AmpC"},
    "Ba04646133_s1": {"Z114", 28, 1, "blaOXA1"},
    "Ba04646139_s1": {"Z116", 29, 1, "OXA-23"},
    "Ba04646138_s1": {"101Y", 28, 1, "blaOXA"},
    "Ba04646137_s1": {"102Y", 26, 1, "ErmA"},
    "Ba04230913_s1": {"103Y", 27, 1, "ermB"},
    "Ba07319994_s1": {"104Y", 30, 1, "ermC"},
    "Ba04230908_s1": {"105Y", 26, 1, "Methicillin 1"},
    "APYMKP3":       {"107Y", 25, 1, "TetB"},
    "Ba07921939_s1": {"107Y", 25, 1, "TetB"},
    "Ba04230915_s1": {"108Y", 25, 1, "tetM"},
    "Ba04646147_s1": {"109Y", 28, 1, "vanA"},
    "Ba04646150_s1": {"110Y", 28, 1, "vanB"},
}
```

