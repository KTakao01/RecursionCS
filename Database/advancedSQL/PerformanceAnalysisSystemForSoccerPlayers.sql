-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1061
SELECT
p.id プレイヤーid,
(p.first_name || ' ' || p.last_name) プレイヤー名,
COUNT(g.id) ゴール数,
SUM(CASE WHEN tb.last_season_rank <= 4 THEN 1 ELSE 0 END) ビッグクラブ相手のゴール数,
t.name 所属チーム,
l.name 所属リーグ
FROM players p
INNER JOIN teams t ON p.team_id = t.id
INNER JOIN leagues l ON l.id = t.league_id
LEFT JOIN goals g ON g.player_id = p.id
JOIN teams tb ON tb.id = g.against_team_id 
GROUP BY p.id
HAVING ゴール数 >= 20 AND ビッグクラブ相手のゴール数 >= 5
ORDER BY ゴール数 DESC;


-- 躓いた点
--　クエリでは SUM(CASE WHEN t.last_season_rank <= 4 THEN 1 ELSE 0 END) という式を使用していますが、この t は players テーブルにリンクされている teams テーブルを指しています。つまり、この式はプレイヤーの所属チームがビッグクラブかどうかを判断しているので、対戦相手チームがビッグクラブかどうかを判定していません。
-- SQLiteではCONCAT関数はサポートされていません。SQLiteで文字列を連結するには、代わりに||演算子を使用します。
-- LEFT JOIN goals g ON g.against_team_id = t.id という結合条件を使うかどうかは、分析したいデータとその文脈によります。ご質問の文脈では、プレイヤーが所属するチームに対してゴールを決めた場合のデータを集計することは意図されていないようです。
--このクエリの目的は、「特定のプレイヤーが他のチームに対してどれだけゴールを決めたか」を分析することです。そのため、goals テーブルの player_id を players テーブルの id と結合することで、特定のプレイヤーが記録したゴールを取得します。また、against_team_id を teams テーブルの id と結合して、そのゴールがどのチームに対して決められたものかを判断します。
--したがって、LEFT JOIN goals g ON g.player_id = p.id はプレイヤーが記録したゴールを抽出し、LEFT JOIN teams bt ON g.against_team_id = bt.id はそのゴールがどのチームに対して記録されたかを判断するために使用します。プレイヤーの所属チーム（t.id）に対してゴールが記録されたかどうかはこの分析では関連がありません。
--しかし、すべてのプレイヤーがゴールを記録しているわけではないため、ゴールがないプレイヤーは結果から除外される可能性があります。そのため、LEFT JOIN を使用して、ゴールが記録されていないプレイヤーも結果に含めることが一般的です。

